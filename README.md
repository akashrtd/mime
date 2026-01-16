# MIME - Modern Browser Automation

**MIME** (Modern Internet Manipulation Engine) is a high-performance browser automation tool built in Go with the Rod library.

## ✨ Features

- 🚀 **High Performance** - Built with Go and Rod (Chrome DevTools Protocol)
- 🎯 **Simple API** - Cleaner, more intuitive than Puppeteer
- 📦 **Single Binary** - No runtime dependencies
- 🔒 **Type Safe** - Full Go type safety
- ⚡ **Fast** - Compiled performance, faster than Node.js

## 🚀 Quick Start

### Installation

#### Build from Source
```bash
git clone https://github.com/akashrtd/mime
cd mime
go build -o mime cmd/mime/main.go
```

### Usage

#### As Go Library

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/akashrtd/mime/pkg/mime"
)

func main() {
    ctx := context.Background()
    
    // Create browser instance
    m, err := mime.New(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer m.Close()
    
    // Navigate and extract data
    if err := m.Navigate("https://example.com"); err != nil {
        log.Fatal(err)
    }
    
    title, err := m.Extract("h1")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Title: %s\n", title)
}
```

## 📖 API Reference

### Creating a Browser Instance

```go
// With default options (headless)
m, err := mime.New(context.Background())

// With custom options
m, err := mime.NewWithOptions(context.Background(), &browser.Options{
    Headless: false,  // Show browser window
    Timeout:  60 * time.Second,
})
```

### Navigation

```go
// Navigate to a URL
err := m.Navigate("https://example.com")

// Get current URL
url := m.URL()

// Get page title
title, err := m.Title()
```

### Element Interaction

```go
// Click an element
err := m.Click("#button-id")

// Type text into an element
err := m.Type("#input-field", "Hello, World!")

// Wait for an element to appear
err := m.WaitFor(".dynamic-content")
```

### Data Extraction

```go
// Extract text from an element
text, err := m.Extract(".title")

// Extract an attribute
href, err := m.ExtractAttr("a.link", "href")

// Get page HTML
html, err := m.HTML()
```

### JavaScript Execution

```go
// Execute JavaScript
result, err := m.Execute("return document.title")
```

### Screenshots

```go
// Capture screenshot (returns base64-encoded PNG)
screenshot, err := m.Screenshot()

// Decode and save
data, _ := base64.StdEncoding.DecodeString(screenshot)
os.WriteFile("screenshot.png", data, 0644)
```

## 📚 Examples

### Web Scraping

```go
m, _ := mime.New(context.Background())
defer m.Close()

m.Navigate("https://news.ycombinator.com")
title, _ := m.Extract(".titleline a")
fmt.Println(title)
```

### Form Automation

```go
m, _ := mime.New(context.Background())
defer m.Close()

m.Navigate("https://example.com/login")
m.Type("#email", "user@example.com")
m.Type("#password", "password123")
m.Click("#login-button")
```

### Screenshots

```go
m, _ := mime.New(context.Background())
defer m.Close()

m.Navigate("https://example.com")
screenshot, _ := m.Screenshot()
// Save or process screenshot
```

## 🏗️ Architecture

```
MIME
├── cmd/mime/          # CLI entry point
├── internal/
│   └── browser/       # Browser automation core (Rod)
├── pkg/mime/         # Public Go API
└── examples/         # Usage examples
```

## 🆚 Why MIME over Puppeteer?

| Feature | MIME | Puppeteer |
|---------|------|-----------|
| **Performance** | ⚡ Compiled Go | JavaScript (Node.js) |
| **Distribution** | 📦 Single binary | 📦 npm + node_modules |
| **Type Safety** | ✅ Go | ✅ TypeScript |
| **API Design** | 🎯 Simpler | 🔧 Complex |
| **Memory Usage** | 💚 Lower | 💛 Higher |
| **Startup Time** | ⚡ Instant | 🐌 Slower |

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

MIT License - see LICENSE file for details

## 🙏 Acknowledgments

- Built with [Rod](https://github.com/go-rod/rod) - A high-level driver for DevTools Protocol

## 🗺️ Roadmap

- [x] Core browser automation
- [x] Simple Go API
- [x] CLI tool
- [ ] MCP (Model Context Protocol) integration
- [ ] TypeScript SDK
- [ ] Advanced selectors
- [ ] Network interception  
- [ ] Browser profiles
