/**
 * MIME - High-level browser automation API
 */
import { MIMEOptions, NavigateResult, ClickResult, TypeResult, ScreenshotResult, HTMLResult, WaitForResult, ScrollResult, HoverResult, ObserveResult, ActResult, CrawlResult, CrawlOptions, MapResult } from './types';
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
    waitFor(selector: string): Promise<WaitForResult>;
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
    scroll(selector?: string, x?: number, y?: number): Promise<ScrollResult>;
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
    hover(selector: string): Promise<HoverResult>;
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
    markdown(fullPage?: boolean): Promise<{
        markdown: string;
        url: string;
    }>;
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
    metadata(): Promise<{
        title: string;
        description: string;
        url: string;
        canonical?: string;
        og?: Record<string, string>;
    }>;
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
    links(): Promise<{
        links: Array<{
            url: string;
            text: string;
        }>;
        count: number;
    }>;
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
    getCookies(): Promise<{
        cookies: Array<{
            name: string;
            value: string;
            domain: string;
            path: string;
        }>;
        count: number;
    }>;
    /**
     * Clear all cookies
     *
     * @example
     * ```typescript
     * await browser.clearCookies();
     * ```
     */
    clearCookies(): Promise<{
        status: string;
    }>;
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
    observe(): Promise<ObserveResult>;
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
    act(instruction: string): Promise<ActResult>;
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
    crawl(url: string, options?: Omit<CrawlOptions, 'url'>): Promise<CrawlResult>;
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
    map(url: string): Promise<MapResult>;
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