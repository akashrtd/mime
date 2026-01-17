/**
 * MIME Browser Automation SDK
 * 
 * TypeScript client for controlling browsers through the Model Context Protocol (MCP).
 * 
 * @example
 * ```typescript
 * import { MIME } from '@mime-browser/sdk';
 * 
 * const browser = await MIME.connect();
 * await browser.navigate('https://example.com');
 * const title = await browser.extract('h1');
 * console.log(title);
 * await browser.close();
 * ```
 */

export * from './types';
export * from './client';
export { MIME } from './mime';
