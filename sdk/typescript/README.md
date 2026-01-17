# @mime-browser/sdk

TypeScript SDK for **MIME** browser automation - Control browsers through the Model Context Protocol (MCP).

## Features

- 🚀 **Simple API** - Intuitive async/await interface
- 📦 **Type Safe** - Full TypeScript support with complete type definitions
- 🔄 **MCP Native** - Built on Model Context Protocol
- ⚡ **High Performance** - Powered by MIME's Go core (27% faster than Puppeteer)

## Installation

```bash
npm install @mime-browser/sdk
```

**Prerequisite**: Install the MIME binary first:
```bash
go install github.com/akashrtd/mime/cmd/mime@latest
```

## Quick Start

```typescript
import { MIME } from '@mime-browser/sdk';

async function main() {
  // Connect to MIME
  const browser = await MIME.connect();

  // Navigate and extract
  await browser.navigate('https://example.com');
  const title = await browser.extract('h1');
  console.log(title);

  // Close when done
  await browser.close();
}

main();
```

## API Reference

### `MIME.connect(options?)`

Connect to a MIME browser automation server.

```typescript
const browser = await MIME.connect({
  binaryPath: '/path/to/mime',  // Default: 'mime'
  timeout: 30000,                // Default: 30000ms
});
```

### Navigation

```typescript
// Navigate to URL
await browser.navigate('https://example.com');

// Get current URL
const url = await browser.url();
```

### Element Interaction

```typescript
// Click element
await browser.click('#button');
await browser.click('.nav-link');

// Type into input
await browser.type('#email', 'user@example.com');
await browser.type('#password', 'secret');

// Wait for element
await browser.waitFor('.content-loaded');
```

### Data Extraction

```typescript
// Extract text
const title = await browser.extract('h1');
const prices = await browser.extract('.price');

// Get page HTML
const { html, url } = await browser.html();

// Execute JavaScript
const count = await browser.execute('return document.links.length');
```

### Screenshots

```typescript
// Take screenshot
const { screenshot } = await browser.screenshot();

// Save to file
import * as fs from 'fs';
const buffer = Buffer.from(screenshot, 'base64');
fs.writeFileSync('page.png', buffer);
```

## Examples

### Web Scraping

```typescript
import { MIME } from '@mime-browser/sdk';

async function scrapeNews() {
  const browser = await MIME.connect();
  
  await browser.navigate('https://news.ycombinator.com');
  
  const articles = await browser.execute(`
    return Array.from(document.querySelectorAll('.titleline a'))
      .slice(0, 10)
      .map(a => ({ title: a.textContent, url: a.href }));
  `);
  
  console.log(articles);
  await browser.close();
}
```

### Form Automation

```typescript
import { MIME } from '@mime-browser/sdk';

async function login() {
  const browser = await MIME.connect();
  
  await browser.navigate('https://example.com/login');
  await browser.type('#email', 'user@example.com');
  await browser.type('#password', 'password123');
  await browser.click('#submit');
  
  await browser.waitFor('.dashboard');
  console.log('Logged in successfully!');
  
  await browser.close();
}
```

## TypeScript Types

Full type definitions are included:

```typescript
import { 
  MIME,
  MIMEOptions,
  NavigateResult,
  ClickResult,
  ExtractResult,
  ScreenshotResult,
} from '@mime-browser/sdk';
```

## License

MIT - see [LICENSE](../../LICENSE)
