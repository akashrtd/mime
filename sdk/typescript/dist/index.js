// src/client.ts
import { spawn } from "child_process";

class MCPClient {
  process = null;
  options;
  requestId = 0;
  pendingRequests = new Map;
  buffer = "";
  initialized = false;
  constructor(options = {}) {
    this.options = {
      binaryPath: options.binaryPath ?? "mime",
      args: options.args ?? ["serve"],
      timeout: options.timeout ?? 30000,
      headless: options.headless ?? true
    };
  }
  async connect() {
    if (this.process) {
      throw new Error("Already connected");
    }
    this.process = spawn(this.options.binaryPath, this.options.args, {
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
  handleData(data) {
    this.buffer += data;
    const lines = this.buffer.split(`
`);
    this.buffer = lines.pop() || "";
    for (const line of lines) {
      if (line.trim()) {
        try {
          const message = JSON.parse(line);
          this.handleMessage(message);
        } catch {}
      }
    }
  }
  handleMessage(message) {
    if (message.id !== undefined) {
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
      this.process.stdin.write(JSON.stringify(request) + `
`);
    });
  }
  sendNotification(method, params) {
    if (!this.process?.stdin) {
      return;
    }
    const notification = {
      jsonrpc: "2.0",
      method,
      params
    };
    this.process.stdin.write(JSON.stringify(notification) + `
`);
  }
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
  async listTools() {
    if (!this.initialized) {
      throw new Error("Not initialized. Call connect() first.");
    }
    const response = await this.sendRequest("tools/list", {});
    return response.tools;
  }
  async disconnect() {
    if (this.process) {
      this.process.kill();
      this.process = null;
      this.initialized = false;
    }
  }
  isConnected() {
    return this.process !== null && this.initialized;
  }
}
// src/mime.ts
class MIME {
  client;
  constructor(client) {
    this.client = client;
  }
  static async connect(options) {
    const client = new MCPClient(options);
    await client.connect();
    return new MIME(client);
  }
  async navigate(url) {
    const response = await this.client.callTool("navigate", { url });
    return this.parseResponse(response);
  }
  async click(selector) {
    const response = await this.client.callTool("click", { selector });
    return this.parseResponse(response);
  }
  async type(selector, text) {
    const response = await this.client.callTool("type", { selector, text });
    return this.parseResponse(response);
  }
  async extract(selector) {
    const response = await this.client.callTool("extract", { selector });
    const result = this.parseResponse(response);
    return result.text;
  }
  async screenshot() {
    const response = await this.client.callTool("screenshot", {});
    return this.parseResponse(response);
  }
  async execute(script) {
    const response = await this.client.callTool("execute", { script });
    const result = this.parseResponse(response);
    return result.result;
  }
  async html() {
    const response = await this.client.callTool("html", {});
    return this.parseResponse(response);
  }
  async url() {
    const { url } = await this.html();
    return url;
  }
  async waitFor(selector, timeout = 30000) {
    const start = Date.now();
    while (Date.now() - start < timeout) {
      try {
        const text = await this.extract(selector);
        if (text !== undefined) {
          return;
        }
      } catch {}
      await this.sleep(100);
    }
    throw new Error(`Timeout waiting for selector: ${selector}`);
  }
  async close() {
    await this.client.disconnect();
  }
  isConnected() {
    return this.client.isConnected();
  }
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
  sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }
}
export {
  MIME,
  MCPClient
};
