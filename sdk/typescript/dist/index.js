"use strict";
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __export = (target, all) => {
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: true });
};
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);

// src/index.ts
var index_exports = {};
__export(index_exports, {
  MCPClient: () => MCPClient,
  MIME: () => MIME
});
module.exports = __toCommonJS(index_exports);

// src/client.ts
var import_child_process = require("child_process");
var MCPClient = class {
  constructor(options = {}) {
    this.process = null;
    this.requestId = 0;
    this.pendingRequests = /* @__PURE__ */ new Map();
    this.buffer = "";
    this.initialized = false;
    this.options = {
      binaryPath: options.binaryPath ?? "mime",
      args: options.args ?? ["serve"],
      timeout: options.timeout ?? 3e4,
      headless: options.headless ?? true
    };
  }
  /**
   * Connect to the MIME MCP server
   */
  async connect() {
    if (this.process) {
      throw new Error("Already connected");
    }
    this.process = (0, import_child_process.spawn)(this.options.binaryPath, this.options.args, {
      stdio: ["pipe", "pipe", "pipe"]
    });
    this.process.stdout?.on("data", (data) => {
      this.handleData(data.toString());
    });
    this.process.stderr?.on("data", (data) => {
      console.error("[MIME]", data.toString());
    });
    this.process.on("error", (err) => {
      throw new Error(`Failed to start MIME: ${err.message}`);
    });
    this.process.on("exit", (code) => {
      this.process = null;
      this.initialized = false;
      if (code !== 0) {
        console.error(`MIME process exited with code ${code}`);
      }
    });
    await this.initialize();
  }
  /**
   * Initialize MCP session with handshake
   */
  async initialize() {
    const response = await this.sendRequest("initialize", {
      protocolVersion: "2024-11-05",
      capabilities: {},
      clientInfo: {
        name: "@mime-browser/sdk",
        version: "0.1.0"
      }
    });
    if (response) {
      this.initialized = true;
      this.sendNotification("notifications/initialized", {});
    }
  }
  /**
   * Handle incoming data from the MCP server
   */
  handleData(data) {
    this.buffer += data;
    const lines = this.buffer.split("\n");
    this.buffer = lines.pop() || "";
    for (const line of lines) {
      if (line.trim()) {
        try {
          const message = JSON.parse(line);
          this.handleMessage(message);
        } catch {
        }
      }
    }
  }
  /**
   * Handle a parsed JSON-RPC message
   */
  handleMessage(message) {
    if (message.id !== void 0) {
      const pending = this.pendingRequests.get(message.id);
      if (pending) {
        this.pendingRequests.delete(message.id);
        if (message.error) {
          pending.reject(new Error(message.error.message));
        } else {
          pending.resolve(message.result);
        }
      }
    }
  }
  /**
   * Send a JSON-RPC request and wait for response
   */
  sendRequest(method, params) {
    return new Promise((resolve, reject) => {
      if (!this.process?.stdin) {
        reject(new Error("Not connected"));
        return;
      }
      const id = ++this.requestId;
      const request = {
        jsonrpc: "2.0",
        id,
        method,
        params
      };
      this.pendingRequests.set(id, { resolve, reject });
      const timeout = setTimeout(() => {
        this.pendingRequests.delete(id);
        reject(new Error(`Request timeout: ${method}`));
      }, this.options.timeout);
      this.pendingRequests.set(id, {
        resolve: (value) => {
          clearTimeout(timeout);
          resolve(value);
        },
        reject: (error) => {
          clearTimeout(timeout);
          reject(error);
        }
      });
      this.process.stdin.write(JSON.stringify(request) + "\n");
    });
  }
  /**
   * Send a notification (no response expected)
   */
  sendNotification(method, params) {
    if (!this.process?.stdin) {
      return;
    }
    const notification = {
      jsonrpc: "2.0",
      method,
      params
    };
    this.process.stdin.write(JSON.stringify(notification) + "\n");
  }
  /**
   * Call an MCP tool
   */
  async callTool(name, args) {
    if (!this.initialized) {
      throw new Error("Not initialized. Call connect() first.");
    }
    const response = await this.sendRequest("tools/call", {
      name,
      arguments: args
    });
    return response;
  }
  /**
   * List available tools
   */
  async listTools() {
    if (!this.initialized) {
      throw new Error("Not initialized. Call connect() first.");
    }
    const response = await this.sendRequest("tools/list", {});
    return response.tools;
  }
  /**
   * Disconnect from the MIME server
   */
  async disconnect() {
    if (this.process) {
      this.process.kill();
      this.process = null;
      this.initialized = false;
    }
  }
  /**
   * Check if connected
   */
  isConnected() {
    return this.process !== null && this.initialized;
  }
};

// src/mime.ts
var MIME = class _MIME {
  constructor(client) {
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
  static async connect(options) {
    const client = new MCPClient(options);
    await client.connect();
    return new _MIME(client);
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
  async navigate(url) {
    const response = await this.client.callTool("navigate", { url });
    return this.parseResponse(response);
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
  async click(selector) {
    const response = await this.client.callTool("click", { selector });
    return this.parseResponse(response);
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
  async type(selector, text) {
    const response = await this.client.callTool("type", { selector, text });
    return this.parseResponse(response);
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
  async extract(selector) {
    const response = await this.client.callTool("extract", { selector });
    const result = this.parseResponse(response);
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
  async screenshot() {
    const response = await this.client.callTool("screenshot", {});
    return this.parseResponse(response);
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
  async execute(script) {
    const response = await this.client.callTool("execute", { script });
    const result = this.parseResponse(response);
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
  async html() {
    const response = await this.client.callTool("html", {});
    return this.parseResponse(response);
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
  async url() {
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
  async waitFor(selector, timeout = 3e4) {
    const start = Date.now();
    while (Date.now() - start < timeout) {
      try {
        const text = await this.extract(selector);
        if (text !== void 0) {
          return;
        }
      } catch {
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
  async close() {
    await this.client.disconnect();
  }
  /**
   * Check if connected to the MIME server
   */
  isConnected() {
    return this.client.isConnected();
  }
  /**
   * Parse MCP tool response
   */
  parseResponse(response) {
    if (response.isError) {
      const errorText = response.content[0]?.text || "Unknown error";
      throw new Error(errorText);
    }
    const textContent = response.content.find((c) => c.type === "text");
    if (!textContent) {
      throw new Error("No text content in response");
    }
    try {
      return JSON.parse(textContent.text);
    } catch {
      return { text: textContent.text };
    }
  }
  /**
   * Sleep for a specified duration
   */
  sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }
};
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  MCPClient,
  MIME
});
