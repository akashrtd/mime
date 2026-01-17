/**
 * Type definitions for MIME Browser Automation SDK
 */

/**
 * Options for connecting to the MIME MCP server
 */
export interface MIMEOptions {
    /**
     * Path to the MIME binary. Defaults to 'mime' (expects it in PATH)
     */
    binaryPath?: string;

    /**
     * Arguments to pass to the MIME binary
     */
    args?: string[];

    /**
     * Timeout in milliseconds for operations. Default: 30000
     */
    timeout?: number;

    /**
     * Whether to run browser in headless mode. Default: true
     */
    headless?: boolean;
}

/**
 * Result of a navigation operation
 */
export interface NavigateResult {
    status: 'success' | 'error';
    url: string;
    error?: string;
}

/**
 * Result of a click operation
 */
export interface ClickResult {
    status: 'success' | 'error';
    selector: string;
    error?: string;
}

/**
 * Result of a type operation
 */
export interface TypeResult {
    status: 'success' | 'error';
    error?: string;
}

/**
 * Result of an extract operation
 */
export interface ExtractResult {
    text: string;
}

/**
 * Result of a screenshot operation
 */
export interface ScreenshotResult {
    /**
     * Base64-encoded PNG image data
     */
    screenshot: string;
    format: 'base64-png';
}

/**
 * Result of a JavaScript execution
 */
export interface ExecuteResult {
    result: unknown;
}

/**
 * Result of getting HTML content
 */
export interface HTMLResult {
    html: string;
    url: string;
}

/**
 * Tool call request for MCP
 */
export interface ToolCall {
    name: string;
    arguments: Record<string, unknown>;
}

/**
 * Tool call response from MCP
 */
export interface ToolResponse {
    content: Array<{
        type: 'text';
        text: string;
    }>;
    isError?: boolean;
}
