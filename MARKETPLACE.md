# MCP Marketplace Submission Guide

## Submission Links

### 1. mcp.so (Primary MCP Directory)
**URL**: https://mcp.so/submit

Submit your GitHub URL: `https://github.com/akashrtd/mime`

### 2. Glama.ai
**URL**: https://glama.ai/mcp/servers

Click "Submit MCP Server" and provide:
- GitHub URL: `https://github.com/akashrtd/mime`
- Description: High-performance browser automation (27% faster than Puppeteer)

### 3. mcpmarket.com (Cline Marketplace)
**URL**: https://mcpmarket.com/submit

Provide:
- Server Name: `mime`
- GitHub: `https://github.com/akashrtd/mime`
- Command: `mime serve`

### 4. Smithery.ai
**URL**: https://smithery.ai/submit

Submit for automatic deployment and hosting.

### 5. Official MCP Registry
**Install publisher CLI**:
```bash
brew install mcp-publisher
```

**Publish**:
```bash
cd /path/to/mime
mcp-publisher init
mcp-publisher login github
mcp-publisher publish
```

## Server Details for Submission

| Field | Value |
|-------|-------|
| **Name** | mime |
| **Description** | High-performance browser automation with MCP. 27% faster than Puppeteer. |
| **GitHub** | https://github.com/akashrtd/mime |
| **Command** | `mime serve` |
| **Language** | Go |
| **License** | MIT |
| **Categories** | Browser, Automation, Web Scraping, Testing |

## Available Tools

- `navigate` - Go to URL
- `click` - Click element
- `type` - Type text
- `extract` - Get element text  
- `screenshot` - Capture page
- `execute` - Run JavaScript
- `html` - Get page HTML
