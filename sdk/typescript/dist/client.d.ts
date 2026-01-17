/**
 * MCP Client for communicating with MIME server
 */
import { MIMEOptions, ToolResponse } from './types';
/**
 * Low-level MCP client for MIME server communication
 */
export declare class MCPClient {
    private process;
    private options;
    private requestId;
    private pendingRequests;
    private buffer;
    private initialized;
    constructor(options?: MIMEOptions);
    /**
     * Connect to the MIME MCP server
     */
    connect(): Promise<void>;
    /**
     * Initialize MCP session with handshake
     */
    private initialize;
    /**
     * Handle incoming data from the MCP server
     */
    private handleData;
    /**
     * Handle a parsed JSON-RPC message
     */
    private handleMessage;
    /**
     * Send a JSON-RPC request and wait for response
     */
    private sendRequest;
    /**
     * Send a notification (no response expected)
     */
    private sendNotification;
    /**
     * Call an MCP tool
     */
    callTool(name: string, args: Record<string, unknown>): Promise<ToolResponse>;
    /**
     * List available tools
     */
    listTools(): Promise<Array<{
        name: string;
        description: string;
    }>>;
    /**
     * Disconnect from the MIME server
     */
    disconnect(): Promise<void>;
    /**
     * Check if connected
     */
    isConnected(): boolean;
}
//# sourceMappingURL=client.d.ts.map