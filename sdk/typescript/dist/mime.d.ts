/**
 * MIME - High-level browser automation API
 */
import { MIMEOptions, NavigateResult, ClickResult, TypeResult, ScreenshotResult, HTMLResult } from './types';
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
export declare class MIME {
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
//# sourceMappingURL=mime.d.ts.map