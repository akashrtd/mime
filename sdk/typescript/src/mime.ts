/**
 * MIME - High-level browser automation API
 */

import { MCPClient } from './client';
import {
    MIMEOptions,
    NavigateResult,
    ClickResult,
    TypeResult,
    ExtractResult,
    ScreenshotResult,
    ExecuteResult,
    HTMLResult,
} from './types';

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
export class MIME {
    private client: MCPClient;

    private constructor(client: MCPClient) {
        this.client = client;
    }

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
    static async connect(options?: MIMEOptions): Promise<MIME> {
        const client = new MCPClient(options);
        await client.connect();
        return new MIME(client);
    }

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
    async navigate(url: string): Promise<NavigateResult> {
        const response = await this.client.callTool('navigate', { url });
        return this.parseResponse<NavigateResult>(response);
    }

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
    async click(selector: string): Promise<ClickResult> {
        const response = await this.client.callTool('click', { selector });
        return this.parseResponse<ClickResult>(response);
    }

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
    async type(selector: string, text: string): Promise<TypeResult> {
        const response = await this.client.callTool('type', { selector, text });
        return this.parseResponse<TypeResult>(response);
    }

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
    async extract(selector: string): Promise<string> {
        const response = await this.client.callTool('extract', { selector });
        const result = this.parseResponse<ExtractResult>(response);
        return result.text;
    }

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
    async screenshot(): Promise<ScreenshotResult> {
        const response = await this.client.callTool('screenshot', {});
        return this.parseResponse<ScreenshotResult>(response);
    }

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
    async execute(script: string): Promise<unknown> {
        const response = await this.client.callTool('execute', { script });
        const result = this.parseResponse<ExecuteResult>(response);
        return result.result;
    }

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
    async html(): Promise<HTMLResult> {
        const response = await this.client.callTool('html', {});
        return this.parseResponse<HTMLResult>(response);
    }

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
    async url(): Promise<string> {
        const { url } = await this.html();
        return url;
    }

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
    async waitFor(selector: string, timeout = 30000): Promise<void> {
        const start = Date.now();
        while (Date.now() - start < timeout) {
            try {
                const text = await this.extract(selector);
                if (text !== undefined) {
                    return;
                }
            } catch {
                // Element not found, retry
            }
            await this.sleep(100);
        }
        throw new Error(`Timeout waiting for selector: ${selector}`);
    }

    /**
     * Close the browser and disconnect
     * 
     * @example
     * ```typescript
     * await browser.close();
     * ```
     */
    async close(): Promise<void> {
        await this.client.disconnect();
    }

    /**
     * Check if connected to the MIME server
     */
    isConnected(): boolean {
        return this.client.isConnected();
    }

    /**
     * Parse MCP tool response
     */
    private parseResponse<T>(response: { content: Array<{ type: string; text: string }>; isError?: boolean }): T {
        if (response.isError) {
            const errorText = response.content[0]?.text || 'Unknown error';
            throw new Error(errorText);
        }

        const textContent = response.content.find(c => c.type === 'text');
        if (!textContent) {
            throw new Error('No text content in response');
        }

        try {
            return JSON.parse(textContent.text) as T;
        } catch {
            // If not JSON, return as-is for text responses
            return { text: textContent.text } as T;
        }
    }

    /**
     * Sleep for a specified duration
     */
    private sleep(ms: number): Promise<void> {
        return new Promise(resolve => setTimeout(resolve, ms));
    }
}
