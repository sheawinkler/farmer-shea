/**
 * @license
 * Copyright 2025 Google LLC
 * SPDX-License-Identifier: Apache-2.0
 */

import {
  vi,
  describe,
  it,
  expect,
  beforeEach,
  afterEach,
  type Mock,
} from 'vitest';
import {
  checkHasEditorType,
  getDiffCommand,
  openDiff,
  allowEditorTypeInSandbox,
  isEditorAvailable,
} from './editor.js';
import { execSync, spawn } from 'child_process';

vi.mock('child_process', () => ({
  execSync: vi.fn(),
  spawn: vi.fn(),
}));

describe('checkHasEditorType', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('should return true for vscode if "code" command exists', () => {
    (execSync as Mock).mockReturnValue(Buffer.from('/usr/bin/code'));
    expect(checkHasEditorType('vscode')).toBe(true);
    const expectedCommand =
      process.platform === 'win32' ? 'where.exe code.cmd' : 'command -v code';
    expect(execSync).toHaveBeenCalledWith(expectedCommand, {
      stdio: 'ignore',
    });
  });

  it('should return false for vscode if "code" command does not exist', () => {
    (execSync as Mock).mockImplementation(() => {
      throw new Error();
    });
    expect(checkHasEditorType('vscode')).toBe(false);
  });

  it('should return true for windsurf if "windsurf" command exists', () => {
    (execSync as Mock).mockReturnValue(Buffer.from('/usr/bin/windsurf'));
    expect(checkHasEditorType('windsurf')).toBe(true);
    expect(execSync).toHaveBeenCalledWith('command -v windsurf', {
      stdio: 'ignore',
    });
  });

  it('should return false for windsurf if "windsurf" command does not exist', () => {
    (execSync as Mock).mockImplementation(() => {
      throw new Error();
    });
    expect(checkHasEditorType('windsurf')).toBe(false);
  });

  it('should return true for cursor if "cursor" command exists', () => {
    (execSync as Mock).mockReturnValue(Buffer.from('/usr/bin/cursor'));
    expect(checkHasEditorType('cursor')).toBe(true);
    expect(execSync).toHaveBeenCalledWith('command -v cursor', {
      stdio: 'ignore',
    });
  });

  it('should return false for cursor if "cursor" command does not exist', () => {
    (execSync as Mock).mockImplementation(() => {
      throw new Error();
    });
    expect(checkHasEditorType('cursor')).toBe(false);
  });

  it('should return true for vim if "vim" command exists', () => {
    (execSync as Mock).mockReturnValue(Buffer.from('/usr/bin/vim'));
    expect(checkHasEditorType('vim')).toBe(true);
    const expectedCommand =
      process.platform === 'win32' ? 'where.exe vim' : 'command -v vim';
    expect(execSync).toHaveBeenCalledWith(expectedCommand, {
      stdio: 'ignore',
    });
  });

  it('should return false for vim if "vim" command does not exist', () => {
    (execSync as Mock).mockImplementation(() => {
      throw new Error();
    });
    expect(checkHasEditorType('vim')).toBe(false);
  });
});

describe('getDiffCommand', () => {
  it('should return the correct command for vscode', () => {
    const command = getDiffCommand('old.txt', 'new.txt', 'vscode');
    expect(command).toEqual({
      command: 'code',
      args: ['--wait', '--diff', 'old.txt', 'new.txt'],
    });
  });

  it('should return the correct command for vim', () => {
    const command = getDiffCommand('old.txt', 'new.txt', 'vim');
    expect(command?.command).toBe('vim');
    expect(command?.args).toContain('old.txt');
    expect(command?.args).toContain('new.txt');
  });

  it('should return null for an unsupported editor', () => {
    // @ts-expect-error Testing unsupported editor
    const command = getDiffCommand('old.txt', 'new.txt', 'foobar');
    expect(command).toBeNull();
  });
});

describe('openDiff', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('should call spawn for vscode', async () => {
    const mockSpawn = {
      on: vi.fn((event, cb) => {
        if (event === 'close') {
          cb(0);
        }
      }),
    };
    (spawn as Mock).mockReturnValue(mockSpawn);
    await openDiff('old.txt', 'new.txt', 'vscode');
    expect(spawn).toHaveBeenCalledWith(
      'code',
      ['--wait', '--diff', 'old.txt', 'new.txt'],
      { stdio: 'inherit' },
    );
    expect(mockSpawn.on).toHaveBeenCalledWith('close', expect.any(Function));
    expect(mockSpawn.on).toHaveBeenCalledWith('error', expect.any(Function));
  });

  it('should call execSync for vim', async () => {
    await openDiff('old.txt', 'new.txt', 'vim');
    expect(execSync).toHaveBeenCalled();
    const command = (execSync as Mock).mock.calls[0][0];
    expect(command).toContain('vim');
    expect(command).toContain('old.txt');
    expect(command).toContain('new.txt');
  });

  it('should handle spawn error for vscode', async () => {
    const mockSpawn = {
      on: vi.fn((event, cb) => {
        if (event === 'error') {
          cb(new Error('spawn error'));
        }
      }),
    };
    (spawn as Mock).mockReturnValue(mockSpawn);
    await expect(openDiff('old.txt', 'new.txt', 'vscode')).rejects.toThrow(
      'spawn error',
    );
  });
});

describe('allowEditorTypeInSandbox', () => {
  afterEach(() => {
    delete process.env.SANDBOX;
  });

  it('should allow vim in sandbox mode', () => {
    process.env.SANDBOX = 'sandbox';
    expect(allowEditorTypeInSandbox('vim')).toBe(true);
  });

  it('should allow vim when not in sandbox mode', () => {
    delete process.env.SANDBOX;
    expect(allowEditorTypeInSandbox('vim')).toBe(true);
  });

  it('should not allow vscode in sandbox mode', () => {
    process.env.SANDBOX = 'sandbox';
    expect(allowEditorTypeInSandbox('vscode')).toBe(false);
  });

  it('should allow vscode when not in sandbox mode', () => {
    delete process.env.SANDBOX;
    expect(allowEditorTypeInSandbox('vscode')).toBe(true);
  });

  it('should not allow windsurf in sandbox mode', () => {
    process.env.SANDBOX = 'sandbox';
    expect(allowEditorTypeInSandbox('windsurf')).toBe(false);
  });

  it('should allow windsurf when not in sandbox mode', () => {
    delete process.env.SANDBOX;
    expect(allowEditorTypeInSandbox('windsurf')).toBe(true);
  });

  it('should not allow cursor in sandbox mode', () => {
    process.env.SANDBOX = 'sandbox';
    expect(allowEditorTypeInSandbox('cursor')).toBe(false);
  });

  it('should allow cursor when not in sandbox mode', () => {
    delete process.env.SANDBOX;
    expect(allowEditorTypeInSandbox('cursor')).toBe(true);
  });
});

describe('isEditorAvailable', () => {
  afterEach(() => {
    delete process.env.SANDBOX;
  });

  it('should return false for undefined editor', () => {
    expect(isEditorAvailable(undefined)).toBe(false);
  });

  it('should return false for empty string editor', () => {
    expect(isEditorAvailable('')).toBe(false);
  });

  it('should return false for invalid editor type', () => {
    expect(isEditorAvailable('invalid-editor')).toBe(false);
  });

  it('should return true for vscode when installed and not in sandbox mode', () => {
    (execSync as Mock).mockReturnValue(Buffer.from('/usr/bin/code'));
    expect(isEditorAvailable('vscode')).toBe(true);
  });

  it('should return false for vscode when not installed and not in sandbox mode', () => {
    (execSync as Mock).mockImplementation(() => {
      throw new Error();
    });
    expect(isEditorAvailable('vscode')).toBe(false);
  });

  it('should return false for vscode when installed and in sandbox mode', () => {
    (execSync as Mock).mockReturnValue(Buffer.from('/usr/bin/code'));
    process.env.SANDBOX = 'sandbox';
    expect(isEditorAvailable('vscode')).toBe(false);
  });
});
