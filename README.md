# MIME

**MIME** (MCP-Integrated Modern Executor) is a high-performance browser automation tool built in Go with native **Model Context Protocol (MCP)** support.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![MCP](https://img.shields.io/badge/MCP-Compatible-green.svg)](https://modelcontextprotocol.io)

## 🏆 Performance

| Metric | MIME | Puppeteer | Improvement |
|--------|------|-----------|-------------|
| **Total** | 1210 ms | 1670 ms | **27% faster** |
| **Navigation** | 187 ms | 808 ms | **77% faster** |
| **Extraction** | 2 ms | 17 ms | **88% faster** |

## ✨ Features

- 🚀 **High Performance** - Built with Go and Rod (Chrome DevTools Protocol)
- 🤖 **MCP Native** - AI agents like Claude can control browsers directly
- 📦 **Single Binary** - No runtime dependencies
- 🎯 **Simple API** - Cleaner than Puppeteer

## 🚀 Installation

```bash
go install github.com/akashrtd/mime/cmd/mime@latest
```

Or build from source:
```bash
git clone https://github.com/akashrtd/mime
cd mime
go build -o mime ./cmd/mime
```

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

| Tool | Description |
|------|-------------|
| `navigate` | Navigate to a URL |
| `click` | Click an element by CSS selector |
| `type` | Type text into an element |
| `extract` | Extract text from an element |
| `screenshot` | Capture screenshot (base64 PNG) |
| `execute` | Run JavaScript on the page |
| `html` | Get page HTML content |

### Example Prompts for Claude

- *"Navigate to https://news.ycombinator.com and extract the top story title"*
- *"Go to example.com, take a screenshot, and describe what you see"*
- *"Click the login button and type my email into the form"*

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

# Show version
mime version
```

## 📊 Benchmarks

See [benchmark/RESULTS.md](benchmark/RESULTS.md) for detailed performance comparison.

```bash
# Run benchmarks
go run benchmark/mime/benchmark.go
```

## 🗺️ Roadmap

- [x] Core browser automation
- [x] Simple Go API  
- [x] CLI tool
- [x] MCP server integration
- [ ] TypeScript SDK
- [ ] Connection pooling
- [ ] Retry logic

## 📄 License

MIT License - see [LICENSE](LICENSE)

## 🙏 Acknowledgments

- [Rod](https://github.com/go-rod/rod) - High-level DevTools Protocol driver
- [MCP](https://modelcontextprotocol.io) - Model Context Protocol by Anthropic
