# Using MIME with MCP

MIME works as an MCP (Model Context Protocol) server, allowing AI assistants like Claude to control browsers directly.

## Quick Setup

### 1. Build MIME

```bash
git clone https://github.com/akashrtd/mime
cd mime
go build -o mime ./cmd/mime
```

### 2. Configure Claude Desktop

Add to your Claude Desktop config:

**macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "mime": {
      "command": "/full/path/to/mime",
      "args": ["serve"]
    }
  }
}
```

### 3. Restart Claude Desktop

## Available Tools

| Tool | Description | Parameters |
|------|-------------|------------|
| `navigate` | Go to a URL | `url` (required) |
| `click` | Click an element | `selector` (required) |
| `type` | Type into an element | `selector`, `text` (required) |
| `extract` | Get text from element | `selector` (required) |
| `screenshot` | Capture page | (none) |
| `execute` | Run JavaScript | `script` (required) |
| `html` | Get page HTML | (none) |

## Example Conversations

### Web Scraping
> "Navigate to https://news.ycombinator.com and tell me the top 3 story titles"

### Form Filling  
> "Go to example.com/login, type 'user@email.com' in the email field, and click submit"

### Visual Testing
> "Navigate to my-app.com and take a screenshot"

### Data Extraction
> "Go to example.com/products and extract all the product names"

## Troubleshooting

**Server not starting?**
- Check the path in config is absolute
- Ensure `mime` has execute permissions: `chmod +x mime`

**Tools not appearing?**
- Restart Claude Desktop after config changes
- Check for typos in config JSON

**Browser issues?**
- MIME auto-downloads Chrome on first run
- Check `~/.cache/rod/browser/` for browser files
