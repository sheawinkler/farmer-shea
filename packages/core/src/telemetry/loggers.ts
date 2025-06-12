/**
 * @license
 * Copyright 2025 Google LLC
 * SPDX-License-Identifier: Apache-2.0
 */

import { logs, LogRecord, LogAttributes } from '@opentelemetry/api-logs';
import { SemanticAttributes } from '@opentelemetry/semantic-conventions';
import { Config } from '../config/config.js';
import {
  EVENT_API_ERROR,
  EVENT_API_REQUEST,
  EVENT_API_RESPONSE,
  EVENT_CLI_CONFIG,
  EVENT_TOOL_CALL,
  EVENT_USER_PROMPT,
  SERVICE_NAME,
} from './constants.js';
import {
  ApiErrorEvent,
  ApiRequestEvent,
  ApiResponseEvent,
  ToolCallEvent,
  UserPromptEvent,
} from './types.js';
import {
  recordApiErrorMetrics,
  recordTokenUsageMetrics,
  recordApiResponseMetrics,
  recordToolCallMetrics,
} from './metrics.js';
import { isTelemetrySdkInitialized } from './sdk.js';
import { ToolConfirmationOutcome } from '../index.js';

const shouldLogUserPrompts = (config: Config): boolean =>
  config.getTelemetryLogUserPromptsEnabled() ?? false;

function getCommonAttributes(config: Config): LogAttributes {
  return {
    'session.id': config.getSessionId(),
  };
}

export enum ToolCallDecision {
  ACCEPT = 'accept',
  REJECT = 'reject',
  MODIFY = 'modify',
}

export function getDecisionFromOutcome(
  outcome: ToolConfirmationOutcome,
): ToolCallDecision {
  switch (outcome) {
    case ToolConfirmationOutcome.ProceedOnce:
    case ToolConfirmationOutcome.ProceedAlways:
    case ToolConfirmationOutcome.ProceedAlwaysServer:
    case ToolConfirmationOutcome.ProceedAlwaysTool:
      return ToolCallDecision.ACCEPT;
    case ToolConfirmationOutcome.ModifyWithEditor:
      return ToolCallDecision.MODIFY;
    case ToolConfirmationOutcome.Cancel:
    default:
      return ToolCallDecision.REJECT;
  }
}

export function logCliConfiguration(config: Config): void {
  if (!isTelemetrySdkInitialized()) return;

  const generatorConfig = config.getContentGeneratorConfig();
  const mcpServers = config.getMcpServers();
  const attributes: LogAttributes = {
    ...getCommonAttributes(config),
    'event.name': EVENT_CLI_CONFIG,
    'event.timestamp': new Date().toISOString(),
    model: config.getModel(),
    embedding_model: config.getEmbeddingModel(),
    sandbox_enabled:
      typeof config.getSandbox() === 'string' ? true : config.getSandbox(),
    core_tools_enabled: (config.getCoreTools() ?? []).join(','),
    approval_mode: config.getApprovalMode(),
    api_key_enabled: !!generatorConfig.apiKey,
    vertex_ai_enabled: !!generatorConfig.vertexai,
    code_assist_enabled: !!generatorConfig.codeAssist,
    log_user_prompts_enabled: config.getTelemetryLogUserPromptsEnabled(),
    file_filtering_respect_git_ignore:
      config.getFileFilteringRespectGitIgnore(),
    file_filtering_allow_build_artifacts:
      config.getFileFilteringAllowBuildArtifacts(),
    debug_mode: config.getDebugMode(),
    mcp_servers: mcpServers ? Object.keys(mcpServers).join(',') : '',
  };
  const logger = logs.getLogger(SERVICE_NAME);
  const logRecord: LogRecord = {
    body: 'CLI configuration loaded.',
    attributes,
  };
  logger.emit(logRecord);
}

export function logUserPrompt(
  config: Config,
  event: Omit<UserPromptEvent, 'event.name' | 'event.timestamp' | 'prompt'> & {
    prompt: string;
  },
): void {
  if (!isTelemetrySdkInitialized()) return;

  const attributes: LogAttributes = {
    ...getCommonAttributes(config),
    'event.name': EVENT_USER_PROMPT,
    'event.timestamp': new Date().toISOString(),
    prompt_length: event.prompt_length,
  };

  if (shouldLogUserPrompts(config)) {
    attributes.prompt = event.prompt;
  }

  const logger = logs.getLogger(SERVICE_NAME);
  const logRecord: LogRecord = {
    body: `User prompt. Length: ${event.prompt_length}`,
    attributes,
  };
  logger.emit(logRecord);
}

export function logToolCall(
  config: Config,
  event: Omit<ToolCallEvent, 'event.name' | 'event.timestamp' | 'decision'>,
  outcome?: ToolConfirmationOutcome,
): void {
  if (!isTelemetrySdkInitialized()) return;

  const decision = outcome ? getDecisionFromOutcome(outcome) : undefined;

  const attributes: LogAttributes = {
    ...getCommonAttributes(config),
    ...event,
    'event.name': EVENT_TOOL_CALL,
    'event.timestamp': new Date().toISOString(),
    function_args: JSON.stringify(event.function_args, null, 2),
    decision,
  };
  if (event.error) {
    attributes['error.message'] = event.error;
    if (event.error_type) {
      attributes['error.type'] = event.error_type;
    }
  }
  const logger = logs.getLogger(SERVICE_NAME);
  const logRecord: LogRecord = {
    body: `Tool call: ${event.function_name}${decision ? `. Decision: ${decision}` : ''}. Success: ${event.success}. Duration: ${event.duration_ms}ms.`,
    attributes,
  };
  logger.emit(logRecord);
  recordToolCallMetrics(
    config,
    event.function_name,
    event.duration_ms,
    event.success,
    decision,
  );
}

export function logApiRequest(
  config: Config,
  event: Omit<ApiRequestEvent, 'event.name' | 'event.timestamp'>,
): void {
  if (!isTelemetrySdkInitialized()) return;
  const attributes: LogAttributes = {
    ...getCommonAttributes(config),
    ...event,
    'event.name': EVENT_API_REQUEST,
    'event.timestamp': new Date().toISOString(),
  };
  const logger = logs.getLogger(SERVICE_NAME);
  const logRecord: LogRecord = {
    body: `API request to ${event.model}. Tokens: ${event.input_token_count}.`,
    attributes,
  };
  logger.emit(logRecord);
  recordTokenUsageMetrics(
    config,
    event.model,
    event.input_token_count,
    'input',
  );
}

export function logApiError(
  config: Config,
  event: Omit<ApiErrorEvent, 'event.name' | 'event.timestamp'>,
): void {
  if (!isTelemetrySdkInitialized()) return;
  const attributes: LogAttributes = {
    ...getCommonAttributes(config),
    ...event,
    'event.name': EVENT_API_ERROR,
    'event.timestamp': new Date().toISOString(),
    ['error.message']: event.error,
  };

  if (event.error_type) {
    attributes['error.type'] = event.error_type;
  }
  if (typeof event.status_code === 'number') {
    attributes[SemanticAttributes.HTTP_STATUS_CODE] = event.status_code;
  }

  const logger = logs.getLogger(SERVICE_NAME);
  const logRecord: LogRecord = {
    body: `API error for ${event.model}. Error: ${event.error}. Duration: ${event.duration_ms}ms.`,
    attributes,
  };
  logger.emit(logRecord);
  recordApiErrorMetrics(
    config,
    event.model,
    event.duration_ms,
    event.status_code,
    event.error_type,
  );
}

export function logApiResponse(
  config: Config,
  event: Omit<ApiResponseEvent, 'event.name' | 'event.timestamp'>,
): void {
  if (!isTelemetrySdkInitialized()) return;
  const attributes: LogAttributes = {
    ...getCommonAttributes(config),
    ...event,
    'event.name': EVENT_API_RESPONSE,
    'event.timestamp': new Date().toISOString(),
  };
  if (event.response_text) {
    attributes.response_text = event.response_text;
  }
  if (event.error) {
    attributes['error.message'] = event.error;
  } else if (event.status_code) {
    if (typeof event.status_code === 'number') {
      attributes[SemanticAttributes.HTTP_STATUS_CODE] = event.status_code;
    }
  }

  const logger = logs.getLogger(SERVICE_NAME);
  const logRecord: LogRecord = {
    body: `API response from ${event.model}. Status: ${event.status_code || 'N/A'}. Duration: ${event.duration_ms}ms.`,
    attributes,
  };
  logger.emit(logRecord);
  recordApiResponseMetrics(
    config,
    event.model,
    event.duration_ms,
    event.status_code,
    event.error,
  );
  recordTokenUsageMetrics(
    config,
    event.model,
    event.output_token_count,
    'output',
  );
  recordTokenUsageMetrics(
    config,
    event.model,
    event.cached_content_token_count,
    'cache',
  );
  recordTokenUsageMetrics(
    config,
    event.model,
    event.thoughts_token_count,
    'thought',
  );
  recordTokenUsageMetrics(config, event.model, event.tool_token_count, 'tool');
}
