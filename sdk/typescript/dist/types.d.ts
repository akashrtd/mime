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
 * Result of a wait_for operation
 */
export interface WaitForResult {
    status: 'success' | 'error';
    error?: string;
}
/**
 * Result of a scroll operation
 */
export interface ScrollResult {
    status: 'success' | 'error';
    error?: string;
}
/**
 * Result of a hover operation
 */
export interface HoverResult {
    status: 'success' | 'error';
    error?: string;
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
/**
 * Result of a markdown extraction
 */
export interface MarkdownResult {
    markdown: string;
    url: string;
}
/**
 * Page metadata
 */
export interface MetadataResult {
    title: string;
    description: string;
    url: string;
    canonical?: string;
    og?: Record<string, string>;
}
/**
 * A hyperlink on the page
 */
export interface Link {
    url: string;
    text: string;
}
/**
 * Result of links extraction
 */
export interface LinksResult {
    links: Link[];
    count: number;
}
/**
 * Cookie information
 */
export interface Cookie {
    name: string;
    value: string;
    domain: string;
    path: string;
}
/**
 * Result of get_cookies
 */
export interface CookiesResult {
    cookies: Cookie[];
    count: number;
}
/**
 * Generic status result
 */
export interface StatusResult {
    status: 'success' | 'error';
    error?: string;
}
/**
 * Page observation result
 */
export interface ObserveResult {
    url: string;
    title: string;
    forms: Array<{
        id: string;
        action: string;
        method: string;
        fields: Array<{
            name: string;
            type: string;
            selector: string;
            placeholder: string;
            required: boolean;
        }>;
        submit?: {
            text: string;
            selector: string;
            type: string;
        };
    }>;
    clickable: Array<{
        text: string;
        selector: string;
        type: string;
    }>;
    inputs: Array<{
        name: string;
        type: string;
        selector: string;
        placeholder: string;
        required: boolean;
    }>;
    content: {
        headings: string[];
        paragraphs: string[];
        image_count: number;
    };
}
/**
 * Act result
 */
export interface ActResult {
    success: boolean;
    action: string;
    target: string;
    message?: string;
}
/**
 * Crawl options
 */
export interface CrawlOptions {
    url: string;
    max_pages?: number;
    max_depth?: number;
    same_domain?: boolean;
    patterns?: string[];
    excludes?: string[];
}
/**
 * Single page result from crawl
 */
export interface PageResult {
    url: string;
    title: string;
    markdown?: string;
    status: string;
}
/**
 * Crawl result
 */
export interface CrawlResult {
    pages: PageResult[];
    total: number;
}
/**
 * Map result
 */
export interface MapResult {
    urls: string[];
    total: number;
}
//# sourceMappingURL=types.d.ts.map