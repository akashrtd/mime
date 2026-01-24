# MIME

**MIME** (MCP-Integrated Modern Executor) is a high-performance browser automation tool built in Go with native **Model Context Protocol (MCP)** support.

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![MCP](https://img.shields.io/badge/MCP-Compatible-green.svg)](https://modelcontextprotocol.io)

## 🏆 Performance

| Metric         | MIME   | Puppeteer | Playwright | vs Puppeteer   | vs Playwright  |
| -------------- | ------ | --------- | ---------- | -------------- | -------------- |
| **Startup**    | 650 ms | 855 ms    | 720 ms     | **24% faster** | **10% faster** |
| **Navigation** | 187 ms | 808 ms    | 450 ms     | **77% faster** | **58% faster** |
| **Screenshot** | 95 ms  | 163 ms    | 122 ms     | **42% faster** | **22% faster** |
| **Extraction** | 2 ms   | 17 ms     | 12 ms      | **88% faster** | **83% faster** |
| **Total**      | 934 ms | 1843 ms   | 1304 ms    | **49% faster** | **28% faster** |

## ✨ Features

- 🚀 **High Performance** - Built with Go and Rod (Chrome DevTools Protocol)
- 🤖 **MCP Native** - AI agents like Claude can control browsers directly
- 📦 **Single Binary** - No runtime dependencies (Node/Python not required)
- 🛡️ **Stealth Mode** - Built-in anti-detection and fingerprint masking
- 🕵️ **Trace Mode** - Record every action with screenshots for debugging
- ⚡ **Auto-Wait** - Intelligent stability checks (waiting for animations/enabled state)

## 🆚 Comparison with Playwright

| Feature          | MIME                           | Playwright                               |
| :--------------- | :----------------------------- | :--------------------------------------- |
| **Architecture** | Single Static Binary (Go)      | Node.js Runtime + NPM Modules            |
| **Use Case**     | **AI Agents & LLMs**           | E2E Testing & QA                         |
| **Installation** | `go install` (Zero Config)     | `npm install` + `npx playwright install` |
| **Performance**  | **High** (Go + Raw CDP)        | Moderate (JS Bridge Overhead)            |
| **Stealth**      | **Native** (Integrated)        | Requires plugins (`stealth-evasions`)    |
| **Security**     | Sandbox-ready, Block `file://` | Full local access by default             |
| **Ecosystem**    | MCP-First (Claude/Cursor)      | Developer-First (VS Code / CI)           |

## 🚀 Installation

> **Note:** This is a private repository. You need repository access to install.

### Option 1: Build from Source (Recommended)

```bash
git clone git@github.com:akashrtd/mime.git
cd mime
make install
```

### Option 2: Install via Go

```bash
# Requires GOPRIVATE to be set for private repo access
export GOPRIVATE=github.com/akashrtd/*
go install github.com/akashrtd/mime/cmd/mime@latest
```

You can also use `make build` to create a binary in `./bin/mime` without installing it system-wide.

## 🤖 MCP Server (AI Integration)

MIME can be used as an MCP server, allowing AI assistants to control browsers.

### Configure with Claude Desktop

Add to `~/Library/Application Support/Claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "mime": {
      "command": "/path/to/mime",
      "args": ["serve"]
    }
  }
}
```

### Available MCP Tools

| Tool            | Description                                                 |
| --------------- | ----------------------------------------------------------- |
| `navigate`      | Navigate to a URL                                           |
| `click`         | Click an element by CSS selector                            |
| `type`          | Type text into an element                                   |
| `extract`       | Extract text from an element                                |
| `screenshot`    | Capture screenshot (PNG/JPEG). Options: `format`, `quality` |
| `execute`       | Run JavaScript on the page                                  |
| `html`          | Get page HTML content                                       |
| `wait_for`      | Wait for element (text/css)                                 |
| `scroll`        | Scroll to element or position                               |
| `hover`         | Hover over an element                                       |
| `markdown`      | Get page content as Markdown                                |
| `metadata`      | Extract page metadata                                       |
| `links`         | Extract all links                                           |
| `get_cookies`   | Get all cookies                                             |
| `clear_cookies` | Clear all cookies                                           |
| `observe`       | Analyze page for AI agents                                  |
| `act`           | Natural language actions                                    |
| `crawl`         | Multi-page crawl                                            |
| `map`           | Site structure mapping                                      |

### Example Prompts for Claude

- _"Navigate to https://news.ycombinator.com and extract the top story title"_
- _"Go to example.com, take a screenshot, and describe what you see"_
- _"Click the login button and type my email into the form"_
- _"Observe the page and tell me what forms are available"_
- _"Crawl https://example.com and save the markdown to a file"_

## 📖 Go Library Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/akashrtd/mime/pkg/mime"
)

func main() {
    m, err := mime.New(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    defer m.Close()

    m.Navigate("https://example.com")
    title, _ := m.Extract("h1")
    fmt.Println(title)
}
```

## 🛠️ CLI Commands

```bash
# Start MCP server
mime serve

# Start with visible browser (debug mode)
mime serve --visible

# Start with stealth mode enabled
mime serve --stealth


# Show version
mime version
```

## 📦 TypeScript SDK

Install the TypeScript SDK for Node.js projects:

```bash
npm install @mime-browser/sdk
```

```typescript
import { MIME } from "@mime-browser/sdk";

const browser = await MIME.connect();
await browser.navigate("https://example.com");
const title = await browser.extract("h1");
await browser.close();
```

See [sdk/typescript/README.md](sdk/typescript/README.md) for full documentation.

## 📊 Benchmarks

See [benchmark/RESULTS.md](benchmark/RESULTS.md) for detailed performance comparison.

```bash
# Run benchmarks
go run benchmark/mime/benchmark.go
```

## 📂 Project Structure

```
mime/
├── cmd/mime/        # CLI entry point
├── pkg/mime/        # Core library code
├── sdk/             # Client SDKs
├── examples/        # Usage examples
├── Makefile         # Build commands
└── go.mod
```

## 🛠️ Development

We use a `Makefile` to standardize development tasks.

```bash
# Show available commands
make help

# Build binary
make build

# Run tests
make test
```

## 🗺️ Roadmap

- [x] Core browser automation
- [x] Simple Go API
- [x] CLI tool
- [x] MCP server integration
- [x] TypeScript SDK
- [x] CI/CD & Automated Releases
- [x] Session Persistence
- [x] Markdown & Metadata
- [x] AI Agent Tools (Observe, Act)
- [x] Multi-Page (Crawl, Map)
- [x] Connection pooling (Library) - Implemented in `pkg/mime/pool.go`
- [x] Retry logic - Implemented in `pkg/mime/retry.go`
- [x] Stealth Mode (Anti-detection)

## 📄 License

MIT License - see [LICENSE](LICENSE)

## 🙏 Acknowledgments

- [Rod](https://github.com/go-rod/rod) - High-level DevTools Protocol driver
- [MCP](https://modelcontextprotocol.io) - Model Context Protocol by Anthropic
