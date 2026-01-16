# Using MIME with MCP

MIME can be used as an MCP (Model Context Protocol) server, allowing AI assistants like Claude to control browsers directly.

## Configuration

### Claude Desktop

Add MIME to your Claude Desktop configuration:

**macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`

**Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

**Linux**: `~/.config/Claude/claude_desktop_config.json`

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

Replace `/path/to/mime` with the actual path to your MIME binary.

## Starting the Server

```bash
mime serve
```

The server will start and listen for MCP connections via stdio.

## Available Tools

Once configured, you can ask Claude to use these browser automation tools:

### navigate
Navigate to a URL:
```
"Can you navigate to https://example.com?"
```

### click
Click an element:
```
"Click the button with selector #submit-button"
```

### type
Type text into an element:
```
"Type 'hello world' into the search box (#search)"
```

### extract
Extract text from an element:
```
"Extract the text from the h1 element"
```

### screenshot
Capture a screenshot:
```
"Take a screenshot of the current page"
```

### execute
Run JavaScript:
```
"Execute this JavaScript: document.title"
```

## Available Resources

You can also ask Claude to read these resources:

- **browser://page** - Current page HTML
- **browser://screenshot** - Page screenshot (base64)
- **browser://url** - Current page URL

## Example Workflows

### Scraping Data
```
"Navigate to https://news.ycombinator.com and extract the titles of the top 5 stories"
```

### Automated Testing
```
"Navigate to https://example.com, click the login button, type 'user@example.com' in the email field, and take a screenshot"
```

### Data Collection
```
"Visit https://example.com/products and extract all product names and prices"
```

## Troubleshooting

### Server Not Starting
- Ensure MIME binary has execute permissions: `chmod +x mime`
- Check the path in your config is absolute and correct

### Tools Not Appearing
- Restart Claude Desktop after config changes
- Check Claude Desktop logs for errors

### Browser Not Launching
- Ensure Chrome/Chromium is installed
- Check system requirements for Rod library
