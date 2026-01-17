/**
 * Type definitions for MIME Browser Automation SDK
 */
/**
 * Options for connecting to the MIME MCP server
 */
interface MIMEOptions {
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
interface NavigateResult {
    status: 'success' | 'error';
    url: string;
    error?: string;
}
/**
 * Result of a click operation
 */
interface ClickResult {
    status: 'success' | 'error';
    selector: string;
    error?: string;
}
/**
 * Result of a type operation
 */
interface TypeResult {
    status: 'success' | 'error';
    error?: string;
}
/**
 * Result of an extract operation
 */
interface ExtractResult {
    text: string;
}
/**
 * Result of a screenshot operation
 */
interface ScreenshotResult {
    /**
     * Base64-encoded PNG image data
     */
    screenshot: string;
    format: 'base64-png';
}
/**
 * Result of a JavaScript execution
 */
interface ExecuteResult {
    result: unknown;
}
/**
 * Result of getting HTML content
 */
interface HTMLResult {
    html: string;
    url: string;
}
/**
 * Tool call request for MCP
 */
interface ToolCall {
    name: string;
    arguments: Record<string, unknown>;
}
/**
 * Tool call response from MCP
 */
interface ToolResponse {
    content: Array<{
        type: 'text';
        text: string;
    }>;
    isError?: boolean;
}

/**
 * MCP Client for communicating with MIME server
 */

/**
 * Low-level MCP client for MIME server communication
 */
declare class MCPClient {
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

/**
 * MIME - High-level browser automation API
 */

/**
 * MIME Browser Automation Client
 *
 * Provides a simple, intuitive API for browser automation through MCP.
 *
 * @example
 * ```typescript
 * const browser = await MIME.connect();
 *
 * await browser.navigate('https://example.com');
 * await browser.click('#login-button');
 * await browser.type('#email', 'user@example.com');
 *
 * const title = await browser.extract('h1');
 * console.log(title);
 *
 * await browser.close();
 * ```
 */
declare class MIME {
    private client;
    private constructor();
    /**
     * Connect to a MIME browser automation server
     *
     * @param options - Connection options
     * @returns Connected MIME instance
     *
     * @example
     * ```typescript
     * // Connect with defaults (expects 'mime' in PATH)
     * const browser = await MIME.connect();
     *
     * // Connect with custom binary path
     * const browser = await MIME.connect({
     *   binaryPath: '/usr/local/bin/mime',
     *   timeout: 60000,
     * });
     * ```
     */
    static connect(options?: MIMEOptions): Promise<MIME>;
    /**
     * Navigate to a URL
     *
     * @param url - The URL to navigate to
     * @returns Navigation result
     *
     * @example
     * ```typescript
     * await browser.navigate('https://example.com');
     * await browser.navigate('https://example.com/login');
     * ```
     */
    navigate(url: string): Promise<NavigateResult>;
    /**
     * Click an element by CSS selector
     *
     * @param selector - CSS selector of the element to click
     * @returns Click result
     *
     * @example
     * ```typescript
     * await browser.click('#submit-button');
     * await browser.click('.nav-link:first-child');
     * await browser.click('[data-testid="login"]');
     * ```
     */
    click(selector: string): Promise<ClickResult>;
    /**
     * Type text into an input element
     *
     * @param selector - CSS selector of the input element
     * @param text - Text to type
     * @returns Type result
     *
     * @example
     * ```typescript
     * await browser.type('#email', 'user@example.com');
     * await browser.type('#password', 'secretpassword');
     * await browser.type('textarea.comment', 'Hello, world!');
     * ```
     */
    type(selector: string, text: string): Promise<TypeResult>;
    /**
     * Extract text content from an element
     *
     * @param selector - CSS selector of the element
     * @returns Extracted text
     *
     * @example
     * ```typescript
     * const title = await browser.extract('h1');
     * const price = await browser.extract('.product-price');
     * const allLinks = await browser.extract('a');
     * ```
     */
    extract(selector: string): Promise<string>;
    /**
     * Capture a screenshot of the current page
     *
     * @returns Base64-encoded PNG screenshot
     *
     * @example
     * ```typescript
     * const screenshot = await browser.screenshot();
     *
     * // Save to file
     * const buffer = Buffer.from(screenshot.screenshot, 'base64');
     * fs.writeFileSync('screenshot.png', buffer);
     * ```
     */
    screenshot(): Promise<ScreenshotResult>;
    /**
     * Execute JavaScript on the page
     *
     * @param script - JavaScript code to execute
     * @returns Execution result
     *
     * @example
     * ```typescript
     * const title = await browser.execute('return document.title');
     * const count = await browser.execute('return document.querySelectorAll("a").length');
     * ```
     */
    execute(script: string): Promise<unknown>;
    /**
     * Get the HTML content of the current page
     *
     * @returns Page HTML and URL
     *
     * @example
     * ```typescript
     * const { html, url } = await browser.html();
     * console.log(`Page at ${url} has ${html.length} characters`);
     * ```
     */
    html(): Promise<HTMLResult>;
    /**
     * Get the current page URL
     *
     * @returns Current URL
     *
     * @example
     * ```typescript
     * const url = await browser.url();
     * console.log(`Currently at: ${url}`);
     * ```
     */
    url(): Promise<string>;
    /**
     * Wait for an element to appear on the page
     *
     * @param selector - CSS selector to wait for
     * @param timeout - Maximum time to wait in milliseconds (default: 30000)
     *
     * @example
     * ```typescript
     * await browser.waitFor('.loading-spinner', { visible: false });
     * await browser.waitFor('.content-loaded');
     * ```
     */
    waitFor(selector: string, timeout?: number): Promise<void>;
    /**
     * Close the browser and disconnect
     *
     * @example
     * ```typescript
     * await browser.close();
     * ```
     */
    close(): Promise<void>;
    /**
     * Check if connected to the MIME server
     */
    isConnected(): boolean;
    /**
     * Parse MCP tool response
     */
    private parseResponse;
    /**
     * Sleep for a specified duration
     */
    private sleep;
}

export { type ClickResult, type ExecuteResult, type ExtractResult, type HTMLResult, MCPClient, MIME, type MIMEOptions, type NavigateResult, type ScreenshotResult, type ToolCall, type ToolResponse, type TypeResult };
