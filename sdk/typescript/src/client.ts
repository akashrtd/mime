/**
 * MCP Client for communicating with MIME server
 */

import { spawn, ChildProcess } from 'child_process';
import { MIMEOptions, ToolResponse } from './types';

/**
 * Low-level MCP client for MIME server communication
 */
export class MCPClient {
    private process: ChildProcess | null = null;
    private options: Required<MIMEOptions>;
    private requestId = 0;
    private pendingRequests = new Map<number, {
        resolve: (value: unknown) => void;
        reject: (error: Error) => void;
    }>();
    private buffer = '';
    private initialized = false;

    constructor(options: MIMEOptions = {}) {
        this.options = {
            binaryPath: options.binaryPath ?? 'mime',
            args: options.args ?? ['serve'],
            timeout: options.timeout ?? 30000,
            headless: options.headless ?? true,
        };
    }

    /**
     * Connect to the MIME MCP server
     */
    async connect(): Promise<void> {
        if (this.process) {
            throw new Error('Already connected');
        }

        this.process = spawn(this.options.binaryPath, this.options.args, {
            stdio: ['pipe', 'pipe', 'pipe'],
        });

        this.process.stdout?.on('data', (data: Buffer) => {
            this.handleData(data.toString());
        });

        this.process.stderr?.on('data', (data: Buffer) => {
            console.error('[MIME]', data.toString());
        });

        this.process.on('error', (err) => {
            throw new Error(`Failed to start MIME: ${err.message}`);
        });

        this.process.on('exit', (code) => {
            this.process = null;
            this.initialized = false;
            if (code !== 0) {
                console.error(`MIME process exited with code ${code}`);
            }
        });

        // Initialize MCP session
        await this.initialize();
    }

    /**
     * Initialize MCP session with handshake
     */
    private async initialize(): Promise<void> {
        const response = await this.sendRequest('initialize', {
            protocolVersion: '2024-11-05',
            capabilities: {},
            clientInfo: {
                name: '@mime-browser/sdk',
                version: '0.1.0',
            },
        });

        if (response) {
            this.initialized = true;
            // Send initialized notification
            this.sendNotification('notifications/initialized', {});
        }
    }

    /**
     * Handle incoming data from the MCP server
     */
    private handleData(data: string): void {
        this.buffer += data;

        // Process complete JSON-RPC messages
        const lines = this.buffer.split('\n');
        this.buffer = lines.pop() || '';

        for (const line of lines) {
            if (line.trim()) {
                try {
                    const message = JSON.parse(line);
                    this.handleMessage(message);
                } catch {
                    // Skip invalid JSON
                }
            }
        }
    }

    /**
     * Handle a parsed JSON-RPC message
     */
    private handleMessage(message: { id?: number; result?: unknown; error?: { message: string } }): void {
        if (message.id !== undefined) {
            const pending = this.pendingRequests.get(message.id);
            if (pending) {
                this.pendingRequests.delete(message.id);
                if (message.error) {
                    pending.reject(new Error(message.error.message));
                } else {
                    pending.resolve(message.result);
                }
            }
        }
    }

    /**
     * Send a JSON-RPC request and wait for response
     */
    private sendRequest(method: string, params: unknown): Promise<unknown> {
        return new Promise((resolve, reject) => {
            if (!this.process?.stdin) {
                reject(new Error('Not connected'));
                return;
            }

            const id = ++this.requestId;
            const request = {
                jsonrpc: '2.0',
                id,
                method,
                params,
            };

            this.pendingRequests.set(id, { resolve, reject });

            const timeout = setTimeout(() => {
                this.pendingRequests.delete(id);
                reject(new Error(`Request timeout: ${method}`));
            }, this.options.timeout);

            this.pendingRequests.set(id, {
                resolve: (value) => {
                    clearTimeout(timeout);
                    resolve(value);
                },
                reject: (error) => {
                    clearTimeout(timeout);
                    reject(error);
                },
            });

            this.process.stdin.write(JSON.stringify(request) + '\n');
        });
    }

    /**
     * Send a notification (no response expected)
     */
    private sendNotification(method: string, params: unknown): void {
        if (!this.process?.stdin) {
            return;
        }

        const notification = {
            jsonrpc: '2.0',
            method,
            params,
        };

        this.process.stdin.write(JSON.stringify(notification) + '\n');
    }

    /**
     * Call an MCP tool
     */
    async callTool(name: string, args: Record<string, unknown>): Promise<ToolResponse> {
        if (!this.initialized) {
            throw new Error('Not initialized. Call connect() first.');
        }

        const response = await this.sendRequest('tools/call', {
            name,
            arguments: args,
        }) as ToolResponse;

        return response;
    }

    /**
     * List available tools
     */
    async listTools(): Promise<Array<{ name: string; description: string }>> {
        if (!this.initialized) {
            throw new Error('Not initialized. Call connect() first.');
        }

        const response = await this.sendRequest('tools/list', {}) as {
            tools: Array<{ name: string; description: string }>;
        };

        return response.tools;
    }

    /**
     * Disconnect from the MIME server
     */
    async disconnect(): Promise<void> {
        if (this.process) {
            this.process.kill();
            this.process = null;
            this.initialized = false;
        }
    }

    /**
     * Check if connected
     */
    isConnected(): boolean {
        return this.process !== null && this.initialized;
    }
}
