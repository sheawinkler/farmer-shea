/**
 * @license
 * Copyright 2025 Google LLC
 * SPDX-License-Identifier: Apache-2.0
 */

import { describe, it, expect, beforeEach, vi, afterEach } from 'vitest';
import type { Mocked } from 'vitest';
import { renderHook, act } from '@testing-library/react';
import { useCompletion } from './useCompletion.js';
import * as fs from 'fs/promises';
import { FileDiscoveryService } from '@gemini-cli/core';

// Mock dependencies
vi.mock('fs/promises');
vi.mock('@gemini-cli/core', async () => {
  const actual = await vi.importActual('@gemini-cli/core');
  return {
    ...actual,
    FileDiscoveryService: vi.fn(),
    isNodeError: vi.fn((error) => error.code === 'ENOENT'),
    escapePath: vi.fn((path) => path),
    unescapePath: vi.fn((path) => path),
    getErrorMessage: vi.fn((error) => error.message),
  };
});

describe('useCompletion git-aware filtering integration', () => {
  let mockFileDiscoveryService: Mocked<FileDiscoveryService>;
  let mockConfig: {
    fileFiltering?: { enabled?: boolean; respectGitignore?: boolean };
  };
  const testCwd = '/test/project';
  const slashCommands = [
    { name: 'help', description: 'Show help', action: vi.fn() },
    { name: 'clear', description: 'Clear screen', action: vi.fn() },
  ];

  beforeEach(() => {
    mockFileDiscoveryService = {
      initialize: vi.fn(),
      shouldIgnoreFile: vi.fn(),
      filterFiles: vi.fn(),
      getIgnoreInfo: vi.fn(() => ({ gitIgnored: [], customIgnored: [] })),
      glob: vi.fn().mockResolvedValue([]),
    };

    mockConfig = {
      getFileFilteringRespectGitIgnore: vi.fn(() => true),
      getFileFilteringAllowBuildArtifacts: vi.fn(() => false),
      getFileService: vi.fn().mockResolvedValue(mockFileDiscoveryService),
    };

    vi.mocked(FileDiscoveryService).mockImplementation(
      () => mockFileDiscoveryService,
    );
    vi.clearAllMocks();
  });

  afterEach(() => {
    vi.restoreAllMocks();
  });

  it('should filter git-ignored directories from @ completions', async () => {
    // Mock fs.readdir to return both regular and git-ignored directories
    vi.mocked(fs.readdir).mockResolvedValue([
      { name: 'src', isDirectory: () => true },
      { name: 'node_modules', isDirectory: () => true },
      { name: 'dist', isDirectory: () => true },
      { name: 'README.md', isDirectory: () => false },
      { name: '.env', isDirectory: () => false },
    ] as Array<{ name: string; isDirectory: () => boolean }>);

    // Mock git ignore service to ignore certain files
    mockFileDiscoveryService.shouldIgnoreFile.mockImplementation(
      (path: string) =>
        path.includes('node_modules') ||
        path.includes('dist') ||
        path.includes('.env'),
    );

    const { result } = renderHook(() =>
      useCompletion('@', testCwd, true, slashCommands, mockConfig),
    );

    // Wait for async operations to complete
    await act(async () => {
      await new Promise((resolve) => setTimeout(resolve, 150)); // Account for debounce
    });

    expect(result.current.suggestions).toHaveLength(2);
    expect(result.current.suggestions).toEqual(
      expect.arrayContaining([
        { label: 'src/', value: 'src/' },
        { label: 'README.md', value: 'README.md' },
      ]),
    );
    expect(result.current.showSuggestions).toBe(true);
  });

  it('should handle recursive search with git-aware filtering', async () => {
    // Mock the recursive file search scenario
    vi.mocked(fs.readdir).mockImplementation(
      async (dirPath: string | Buffer | URL) => {
        if (dirPath === testCwd) {
          return [
            { name: 'src', isDirectory: () => true },
            { name: 'node_modules', isDirectory: () => true },
            { name: 'temp', isDirectory: () => true },
          ] as Array<{ name: string; isDirectory: () => boolean }>;
        }
        if (dirPath.endsWith('/src')) {
          return [
            { name: 'index.ts', isDirectory: () => false },
            { name: 'components', isDirectory: () => true },
          ] as Array<{ name: string; isDirectory: () => boolean }>;
        }
        if (dirPath.endsWith('/temp')) {
          return [{ name: 'temp.log', isDirectory: () => false }] as Array<{
            name: string;
            isDirectory: () => boolean;
          }>;
        }
        return [] as Array<{ name: string; isDirectory: () => boolean }>;
      },
    );

    // Mock git ignore service
    mockFileDiscoveryService.shouldIgnoreFile.mockImplementation(
      (path: string) => path.includes('node_modules') || path.includes('temp'),
    );

    const { result } = renderHook(() =>
      useCompletion('@t', testCwd, true, slashCommands, mockConfig),
    );

    // Wait for async operations to complete
    await act(async () => {
      await new Promise((resolve) => setTimeout(resolve, 150));
    });

    // Should not include anything from node_modules or dist
    const suggestionLabels = result.current.suggestions.map((s) => s.label);
    expect(suggestionLabels).not.toContain('temp/');
    expect(suggestionLabels.some((l) => l.includes('node_modules'))).toBe(
      false,
    );
  });

  it('should work without config (fallback behavior)', async () => {
    vi.mocked(fs.readdir).mockResolvedValue([
      { name: 'src', isDirectory: () => true },
      { name: 'node_modules', isDirectory: () => true },
      { name: 'README.md', isDirectory: () => false },
    ] as Array<{ name: string; isDirectory: () => boolean }>);

    const { result } = renderHook(() =>
      useCompletion('@', testCwd, true, slashCommands, undefined),
    );

    await act(async () => {
      await new Promise((resolve) => setTimeout(resolve, 150));
    });

    // Without config, should include all files
    expect(result.current.suggestions).toHaveLength(3);
    expect(result.current.suggestions).toEqual(
      expect.arrayContaining([
        { label: 'src/', value: 'src/' },
        { label: 'node_modules/', value: 'node_modules/' },
        { label: 'README.md', value: 'README.md' },
      ]),
    );
  });

  it('should handle git discovery service initialization failure gracefully', async () => {
    mockFileDiscoveryService.initialize.mockRejectedValue(
      new Error('Git not found'),
    );

    vi.mocked(fs.readdir).mockResolvedValue([
      { name: 'src', isDirectory: () => true },
      { name: 'README.md', isDirectory: () => false },
    ] as Array<{ name: string; isDirectory: () => boolean }>);

    const consoleSpy = vi.spyOn(console, 'warn').mockImplementation(() => {});

    const { result } = renderHook(() =>
      useCompletion('@', testCwd, true, slashCommands, mockConfig),
    );

    await act(async () => {
      await new Promise((resolve) => setTimeout(resolve, 150));
    });

    // Since we use centralized service, initialization errors are handled at config level
    // This test should verify graceful fallback behavior
    expect(result.current.suggestions.length).toBeGreaterThanOrEqual(0);
    // Should still show completions even if git discovery fails
    expect(result.current.suggestions.length).toBeGreaterThan(0);

    consoleSpy.mockRestore();
  });

  it('should handle directory-specific completions with git filtering', async () => {
    vi.mocked(fs.readdir).mockResolvedValue([
      { name: 'component.tsx', isDirectory: () => false },
      { name: 'temp.log', isDirectory: () => false },
      { name: 'index.ts', isDirectory: () => false },
    ] as Array<{ name: string; isDirectory: () => boolean }>);

    mockFileDiscoveryService.shouldIgnoreFile.mockImplementation(
      (path: string) => path.includes('.log'),
    );

    const { result } = renderHook(() =>
      useCompletion('@src/comp', testCwd, true, slashCommands, mockConfig),
    );

    await act(async () => {
      await new Promise((resolve) => setTimeout(resolve, 150));
    });

    // Should filter out .log files but include matching .tsx files
    expect(result.current.suggestions).toEqual([
      { label: 'component.tsx', value: 'component.tsx' },
    ]);
  });

  it('should use glob for top-level @ completions when available', async () => {
    const globResults = [`${testCwd}/src/index.ts`, `${testCwd}/README.md`];
    mockFileDiscoveryService.glob.mockResolvedValue(globResults);

    const { result } = renderHook(() =>
      useCompletion('@s', testCwd, true, slashCommands, mockConfig),
    );

    await act(async () => {
      await new Promise((resolve) => setTimeout(resolve, 150));
    });

    expect(mockFileDiscoveryService.glob).toHaveBeenCalledWith('**/s*', {
      cwd: testCwd,
      dot: false,
    });
    expect(fs.readdir).not.toHaveBeenCalled(); // Ensure glob is used instead of readdir
    expect(result.current.suggestions).toEqual([
      { label: 'README.md', value: 'README.md' },
      { label: 'src/index.ts', value: 'src/index.ts' },
    ]);
  });

  it('should include dotfiles in glob search when input starts with a dot', async () => {
    const globResults = [
      `${testCwd}/.env`,
      `${testCwd}/.gitignore`,
      `${testCwd}/src/index.ts`,
    ];
    mockFileDiscoveryService.glob.mockResolvedValue(globResults);

    const { result } = renderHook(() =>
      useCompletion('@.', testCwd, true, slashCommands, mockConfig),
    );

    await act(async () => {
      await new Promise((resolve) => setTimeout(resolve, 150));
    });

    expect(mockFileDiscoveryService.glob).toHaveBeenCalledWith('**/.*', {
      cwd: testCwd,
      dot: true,
    });
    expect(fs.readdir).not.toHaveBeenCalled();
    expect(result.current.suggestions).toEqual([
      { label: '.env', value: '.env' },
      { label: '.gitignore', value: '.gitignore' },
      { label: 'src/index.ts', value: 'src/index.ts' },
    ]);
  });
});
