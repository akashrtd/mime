package mime

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// MCPServer represents an MCP server for browser automation
type MCPServer struct {
	server  *mcp.Server
	browser *Browser
	ctx     context.Context
}

// Tool input/output types

// NavigateInput for navigate tool
type NavigateInput struct {
	URL string `json:"url"`
}

// NavigateOutput for navigate tool
type NavigateOutput struct {
	Status string `json:"status"`
	URL    string `json:"url"`
}

// ClickInput for click tool
type ClickInput struct {
	Selector string `json:"selector"`
}

// ClickOutput for click tool
type ClickOutput struct {
	Status   string `json:"status"`
	Selector string `json:"selector"`
}

// TypeInput for type tool
type TypeInput struct {
	Selector string `json:"selector"`
	Text     string `json:"text"`
}

// TypeOutput for type tool
type TypeOutput struct {
	Status string `json:"status"`
}

// ExtractInput for extract tool
type ExtractInput struct {
	Selector string `json:"selector"`
}

// ExtractOutput for extract tool
type ExtractOutput struct {
	Text string `json:"text"`
}

// ScreenshotInput for screenshot tool (empty - no params needed)
// ScreenshotInput for screenshot tool
type ScreenshotInput struct {
	Format  string `json:"format,omitempty"`  // "png" (default) or "jpeg"
	Quality int    `json:"quality,omitempty"` // 0-100, only for jpeg
}

// ScreenshotOutput for screenshot tool
type ScreenshotOutput struct {
	Screenshot string `json:"screenshot"`
	Format     string `json:"format"`
}

// ExecuteInput for execute tool
type ExecuteInput struct {
	Script string `json:"script"`
}

// ExecuteOutput for execute tool
type ExecuteOutput struct {
	Result interface{} `json:"result"`
}

// HTMLInput for html tool (empty - no params needed)
type HTMLInput struct{}

// HTMLOutput for html tool
type HTMLOutput struct {
	HTML string `json:"html"`
	URL  string `json:"url"`
}

// WaitForOutput for wait_for tool
type WaitForOutput struct {
	Status string `json:"status"`
}

// ScrollOutput for scroll tool
type ScrollOutput struct {
	Status string `json:"status"`
}

// HoverOutput for hover tool
type HoverOutput struct {
	Status string `json:"status"`
}

// NewMCPServer creates a new MCP server for browser automation
func NewMCPServer(ctx context.Context, opts *BrowserOptions) (*MCPServer, error) {
	implementation := &mcp.Implementation{
		Name:    "mime",
		Version: "0.1.0",
	}

	srv := mcp.NewServer(implementation, nil)

	s := &MCPServer{
		server: srv,
		ctx:    ctx,
	}

	// Initialize browser with options
	br, err := NewBrowser(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create browser: %w", err)
	}
	s.browser = br

	// Register tools
	s.registerTools()

	return s, nil
}

// registerTools registers all MCP tools
func (s *MCPServer) registerTools() {
	// Navigate tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "navigate",
			Description: "Navigate the browser to a URL",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input NavigateInput) (*mcp.CallToolResult, NavigateOutput, error) {
			if err := s.browser.Navigate(input.URL); err != nil {
				return nil, NavigateOutput{}, fmt.Errorf("navigation failed: %w", err)
			}
			return nil, NavigateOutput{Status: "success", URL: input.URL}, nil
		},
	)

	// Click tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "click",
			Description: "Click an element on the page by CSS selector (supports 'text=Login')",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input ClickInput) (*mcp.CallToolResult, ClickOutput, error) {
			if err := s.browser.Click(input.Selector); err != nil {
				return nil, ClickOutput{}, fmt.Errorf("click failed: %w", err)
			}
			return nil, ClickOutput{Status: "success", Selector: input.Selector}, nil
		},
	)

	// Type tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "type",
			Description: "Type text into an element on the page",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input TypeInput) (*mcp.CallToolResult, TypeOutput, error) {
			if err := s.browser.Type(input.Selector, input.Text); err != nil {
				return nil, TypeOutput{}, fmt.Errorf("type failed: %w", err)
			}
			return nil, TypeOutput{Status: "success"}, nil
		},
	)

	// Extract tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "extract",
			Description: "Extract text content from an element on the page",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input ExtractInput) (*mcp.CallToolResult, ExtractOutput, error) {
			text, err := s.browser.Extract(input.Selector)
			if err != nil {
				return nil, ExtractOutput{}, fmt.Errorf("extract failed: %w", err)
			}
			return nil, ExtractOutput{Text: text}, nil
		},
	)

	// Screenshot tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "screenshot",
			Description: "Capture a screenshot of the current page (returns base64-encoded PNG)",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input ScreenshotInput) (*mcp.CallToolResult, ScreenshotOutput, error) {
			data, err := s.browser.Screenshot(input.Format, input.Quality)
			if err != nil {
				return nil, ScreenshotOutput{}, fmt.Errorf("screenshot failed: %w", err)
			}
			format := "base64-png"
			if input.Format == "jpeg" || input.Format == "jpg" {
				format = "base64-jpeg"
			}
			return nil, ScreenshotOutput{Screenshot: data, Format: format}, nil
		},
	)

	// Execute JavaScript tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "execute",
			Description: "Execute JavaScript code on the current page",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input ExecuteInput) (*mcp.CallToolResult, ExecuteOutput, error) {
			result, err := s.browser.Execute(input.Script)
			if err != nil {
				return nil, ExecuteOutput{}, fmt.Errorf("execute failed: %w", err)
			}
			return nil, ExecuteOutput{Result: result}, nil
		},
	)

	// Get HTML tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "html",
			Description: "Get the HTML content of the current page",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input HTMLInput) (*mcp.CallToolResult, HTMLOutput, error) {
			html, err := s.browser.HTML()
			if err != nil {
				return nil, HTMLOutput{}, fmt.Errorf("html failed: %w", err)
			}
			return nil, HTMLOutput{HTML: html, URL: s.browser.URL()}, nil
		},
	)

	// NEW TOOLS

	// WaitFor tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "wait_for",
			Description: "Wait for an element to appear (supports 'text=Query')",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input WaitForArgs) (*mcp.CallToolResult, WaitForOutput, error) {
			if err := s.browser.WaitFor(input.Selector); err != nil {
				return nil, WaitForOutput{}, fmt.Errorf("wait failed: %w", err)
			}
			return nil, WaitForOutput{Status: "success"}, nil
		},
	)

	// Scroll tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "scroll",
			Description: "Scroll the window or an element",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input ScrollArgs) (*mcp.CallToolResult, ScrollOutput, error) {
			if err := s.browser.Scroll(input.Selector, input.X, input.Y); err != nil {
				return nil, ScrollOutput{}, fmt.Errorf("scroll failed: %w", err)
			}
			return nil, ScrollOutput{Status: "success"}, nil
		},
	)

	// Hover tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "hover",
			Description: "Hover over an element",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input HoverArgs) (*mcp.CallToolResult, HoverOutput, error) {
			if err := s.browser.Hover(input.Selector); err != nil {
				return nil, HoverOutput{}, fmt.Errorf("hover failed: %w", err)
			}
			return nil, HoverOutput{Status: "success"}, nil
		},
	)

	// Markdown tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "markdown",
			Description: "Get page content as clean markdown (main content only, ideal for LLMs)",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input MarkdownInput) (*mcp.CallToolResult, MarkdownOutput, error) {
			opts := &MarkdownOptions{OnlyMainContent: !input.FullPage}
			md, err := s.browser.Markdown(opts)
			if err != nil {
				return nil, MarkdownOutput{}, fmt.Errorf("markdown failed: %w", err)
			}
			return nil, MarkdownOutput{Markdown: md, URL: s.browser.URL()}, nil
		},
	)

	// Metadata tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "metadata",
			Description: "Extract page metadata (title, description, og tags)",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, MetadataOutput, error) {
			meta, err := s.browser.Metadata()
			if err != nil {
				return nil, MetadataOutput{}, fmt.Errorf("metadata failed: %w", err)
			}
			return nil, MetadataOutput{
				Title:       meta.Title,
				Description: meta.Description,
				URL:         meta.URL,
				Canonical:   meta.Canonical,
				OG:          meta.OG,
			}, nil
		},
	)

	// Links tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "links",
			Description: "Extract all links from the page",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, LinksOutput, error) {
			links, err := s.browser.Links()
			if err != nil {
				return nil, LinksOutput{}, fmt.Errorf("links failed: %w", err)
			}
			return nil, LinksOutput{Links: links, Count: len(links)}, nil
		},
	)

	// Get cookies tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "get_cookies",
			Description: "Get all cookies for the current page",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, CookiesOutput, error) {
			cookies, err := s.browser.GetCookies()
			if err != nil {
				return nil, CookiesOutput{}, fmt.Errorf("get_cookies failed: %w", err)
			}
			// Convert to simpler format
			simpleCookies := make([]CookieInfo, 0, len(cookies))
			for _, c := range cookies {
				simpleCookies = append(simpleCookies, CookieInfo{
					Name:   c.Name,
					Value:  c.Value,
					Domain: c.Domain,
					Path:   c.Path,
				})
			}
			return nil, CookiesOutput{Cookies: simpleCookies, Count: len(simpleCookies)}, nil
		},
	)

	// Clear cookies tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "clear_cookies",
			Description: "Clear all cookies",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, StatusOutput, error) {
			if err := s.browser.ClearCookies(); err != nil {
				return nil, StatusOutput{}, fmt.Errorf("clear_cookies failed: %w", err)
			}
			return nil, StatusOutput{Status: "success"}, nil
		},
	)

	// Observe tool - AI page understanding
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "observe",
			Description: "Analyze page structure for AI agents. Returns forms, clickable elements, inputs, and content summary.",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, ObserveOutput, error) {
			obs, err := s.browser.Observe()
			if err != nil {
				return nil, ObserveOutput{}, fmt.Errorf("observe failed: %w", err)
			}
			return nil, ObserveOutput{
				URL:       obs.URL,
				Title:     obs.Title,
				Forms:     obs.Forms,
				Clickable: obs.Clickable,
				Inputs:    obs.Inputs,
				Content:   obs.Content,
			}, nil
		},
	)

	// Act tool - Natural language actions
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "act",
			Description: "Perform action from natural language. Examples: 'click login button', 'type email into username field'",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input ActInput) (*mcp.CallToolResult, ActOutput, error) {
			result, err := s.browser.Act(input.Instruction)
			if err != nil {
				return nil, ActOutput{}, fmt.Errorf("act failed: %w", err)
			}
			return nil, ActOutput{
				Success: result.Success,
				Action:  result.Action,
				Target:  result.Target,
				Message: result.Message,
			}, nil
		},
	)

	// Crawl tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "crawl",
			Description: "Crawl multiple pages starting from a URL and return markdown content",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input CrawlOptions) (*mcp.CallToolResult, CrawlResult, error) {
			result, err := s.browser.Crawl(input.URL, &input)
			if err != nil {
				return nil, CrawlResult{}, fmt.Errorf("crawl failed: %w", err)
			}
			return nil, *result, nil
		},
	)

	// Map tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "map",
			Description: "Discover all specific pages on a website (sitemap generation)",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input struct {
			URL string `json:"url"`
		}) (*mcp.CallToolResult, MapResult, error) {
			result, err := s.browser.Map(input.URL)
			if err != nil {
				return nil, MapResult{}, fmt.Errorf("map failed: %w", err)
			}
			return nil, *result, nil
		},
	)
}

// Run starts the MCP server with the given transport
func (s *MCPServer) Run(transport mcp.Transport) error {
	return s.server.Run(s.ctx, transport)
}

// Close closes the browser and cleans up resources
func (s *MCPServer) Close() error {
	if s.browser != nil {
		return s.browser.Close()
	}
	return nil
}
