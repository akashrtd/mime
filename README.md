# MIME

**MIME** (Modern Internet Manipulation Engine) is a high-performance browser automation library built in Go, designed as a faster alternative to Puppeteer.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## 🏆 Performance

| Metric | MIME | Puppeteer | Improvement |
|--------|------|-----------|-------------|
| **Total** | 1210 ms | 1670 ms | **27% faster** |
| **Navigation** | 187 ms | 808 ms | **77% faster** |
| **Extraction** | 2 ms | 17 ms | **88% faster** |

## ✨ Features

- 🚀 **High Performance** - Built with Go and Rod (Chrome DevTools Protocol)
- 📦 **Single Binary** - No runtime dependencies
- 🎯 **Simple API** - Cleaner than Puppeteer
- ⚡ **Fast Startup** - Compiled, not interpreted

## 🚀 Installation

```bash
go get github.com/akashrtd/mime
```

Or build from source:
```bash
git clone https://github.com/akashrtd/mime
cd mime
go build -o mime cmd/mime/main.go
```

## 📖 Usage

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
    
    // Navigate
    m.Navigate("https://example.com")
    
    // Extract data
    title, _ := m.Extract("h1")
    fmt.Println(title)
    
    // Take screenshot
    screenshot, _ := m.Screenshot()
}
```

## 🛠️ API

| Method | Description |
|--------|-------------|
| `Navigate(url)` | Navigate to URL |
| `Click(selector)` | Click element |
| `Type(selector, text)` | Type into element |
| `Extract(selector)` | Extract text |
| `ExtractAttr(selector, attr)` | Extract attribute |
| `Screenshot()` | Capture screenshot (base64 PNG) |
| `HTML()` | Get page HTML |
| `Execute(script)` | Run JavaScript |
| `WaitFor(selector)` | Wait for element |

## 📊 Benchmarks

Run benchmarks yourself:
```bash
# MIME benchmark
go run benchmark/mime/benchmark.go

# Puppeteer benchmark
cd benchmark/puppeteer && npm install && node benchmark.js
```

See [benchmark/RESULTS.md](benchmark/RESULTS.md) for detailed analysis.

## 🗺️ Roadmap

- [x] Core browser automation
- [x] Simple Go API  
- [x] CLI tool
- [ ] MCP (Model Context Protocol) integration
- [ ] TypeScript SDK
- [ ] Connection pooling
- [ ] Retry logic

## 📄 License

MIT License - see [LICENSE](LICENSE)

## 🙏 Acknowledgments

Built with [Rod](https://github.com/go-rod/rod) - High-level DevTools Protocol driver
