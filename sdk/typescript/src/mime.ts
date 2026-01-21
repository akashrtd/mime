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
    WaitForResult,
    ScrollResult,
    HoverResult,
    MarkdownResult,
    MetadataResult,
    Link,
    LinksResult,
    Cookie,
    CookiesResult,
    StatusResult,
    ObserveResult,
    ActResult,
    CrawlResult,
    CrawlOptions,
    MapResult
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
    async waitFor(selector: string): Promise<WaitForResult> {
        const response = await this.client.callTool('wait_for', { selector });
        return this.parseResponse<WaitForResult>(response);
    }

    /**
     * Scroll the window or an element
     * 
     * @param selector - Optional CSS selector to scroll into view
     * @param x - Optional pixels to scroll horizontally
     * @param y - Optional pixels to scroll vertically
     * 
     * @example
     * ```typescript
     * await browser.scroll('#footer');
     * await browser.scroll(undefined, 0, 500); // Scroll down 500px
     * ```
     */
    async scroll(selector?: string, x?: number, y?: number): Promise<ScrollResult> {
        const response = await this.client.callTool('scroll', { selector, x, y });
        return this.parseResponse<ScrollResult>(response);
    }

    /**
     * Hover over an element
     * 
     * @param selector - CSS selector of the element to hover
     * 
     * @example
     * ```typescript
     * await browser.hover('.dropdown-menu');
     * ```
     */
    async hover(selector: string): Promise<HoverResult> {
        const response = await this.client.callTool('hover', { selector });
        return this.parseResponse<HoverResult>(response);
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
     * Get page content as clean markdown (ideal for LLMs)
     * 
     * @param fullPage - If true, includes full page; otherwise main content only
     * @returns Markdown content and URL
     * 
     * @example
     * ```typescript
     * const { markdown } = await browser.markdown();
     * console.log(markdown);
     * ```
     */
    async markdown(fullPage = false): Promise<{ markdown: string; url: string }> {
        const response = await this.client.callTool('markdown', { full_page: fullPage });
        return this.parseResponse<{ markdown: string; url: string }>(response);
    }

    /**
     * Extract page metadata (title, description, og tags)
     * 
     * @returns Page metadata
     * 
     * @example
     * ```typescript
     * const meta = await browser.metadata();
     * console.log(meta.title, meta.description);
     * ```
     */
    async metadata(): Promise<{
        title: string;
        description: string;
        url: string;
        canonical?: string;
        og?: Record<string, string>;
    }> {
        const response = await this.client.callTool('metadata', {});
        return this.parseResponse(response);
    }

    /**
     * Extract all links from the page
     * 
     * @returns Array of links with URL and text
     * 
     * @example
     * ```typescript
     * const { links } = await browser.links();
     * links.forEach(link => console.log(link.url, link.text));
     * ```
     */
    async links(): Promise<{ links: Array<{ url: string; text: string }>; count: number }> {
        const response = await this.client.callTool('links', {});
        return this.parseResponse(response);
    }

    /**
     * Get all cookies for the current page
     * 
     * @returns Array of cookies
     * 
     * @example
     * ```typescript
     * const { cookies } = await browser.getCookies();
     * console.log(`Found ${cookies.length} cookies`);
     * ```
     */
    async getCookies(): Promise<{
        cookies: Array<{ name: string; value: string; domain: string; path: string }>;
        count: number;
    }> {
        const response = await this.client.callTool('get_cookies', {});
        return this.parseResponse(response);
    }

    /**
     * Clear all cookies
     * 
     * @example
     * ```typescript
     * await browser.clearCookies();
     * ```
     */
    async clearCookies(): Promise<{ status: string }> {
        const response = await this.client.callTool('clear_cookies', {});
        return this.parseResponse(response);
    }

    /**
     * Analyze page structure for AI understanding
     * 
     * Returns forms, clickable elements, inputs, and content summary.
     * This is the key tool for AI agents to understand what's on a page.
     * 
     * @example
     * ```typescript
     * const obs = await browser.observe();
     * console.log('Forms:', obs.forms.length);
     * console.log('Clickable:', obs.clickable.map(c => c.text));
     * ```
     */
    async observe(): Promise<ObserveResult> {
        const response = await this.client.callTool('observe', {});
        return this.parseResponse(response);
    }

    /**
     * Perform action from natural language instruction
     * 
     * @param instruction - Natural language instruction like "click login button"
     * @returns Result with success status and action taken
     * 
     * @example
     * ```typescript
     * await browser.act('click the login button');
     * await browser.act('type hello into the search field');
     * await browser.act('scroll to pricing');
     * ```
     */
    async act(instruction: string): Promise<ActResult> {
        const response = await this.client.callTool('act', { instruction });
        return this.parseResponse(response);
    }

    /**
     * Crawl multiple pages starting from a URL
     * 
     * @param url - Starting URL
     * @param options - Crawl options (max_pages, max_depth, etc.)
     * @returns Crawl result with content of all visited pages
     * 
     * @example
     * ```typescript
     * const result = await browser.crawl('https://example.com', { max_pages: 5 });
     * console.log(`Crawled ${result.total} pages`);
     * ```
     */
    async crawl(url: string, options: Omit<CrawlOptions, 'url'> = {}): Promise<CrawlResult> {
        const response = await this.client.callTool('crawl', { url, ...options });
        return this.parseResponse(response);
    }

    /**
     * Map a website structure (discover URLs)
     * 
     * @param url - URL to map
     * @returns List of discovered URLs
     * 
     * @example
     * ```typescript
     * const result = await browser.map('https://example.com');
     * console.log(`Found ${result.total} links`);
     * ```
     */
    async map(url: string): Promise<MapResult> {
        const response = await this.client.callTool('map', { url });
        return this.parseResponse(response);
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
