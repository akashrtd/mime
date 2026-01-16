package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/akashrtd/mime/internal/browser"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Server represents an MCP server for browser automation
type Server struct {
	server  *mcp.Server
	browser *browser.Browser
	ctx     context.Context
}

// Tool input/output types

// NavigateInput for navigate tool
type NavigateInput struct {
	URL string `json:"url" jsonschema:"required,description=URL to navigate to"`
}

// NavigateOutput for navigate tool
type NavigateOutput struct {
	Status string `json:"status"`
	URL    string `json:"url"`
}

// ClickInput for click tool
type ClickInput struct {
	Selector string `json:"selector" jsonschema:"required,description=CSS selector of element to click"`
}

// ClickOutput for click tool
type ClickOutput struct {
	Status   string `json:"status"`
	Selector string `json:"selector"`
}

// TypeInput for type tool
type TypeInput struct {
	Selector string `json:"selector" jsonschema:"required,description=CSS selector of element to type into"`
	Text     string `json:"text" jsonschema:"required,description=Text to type into the element"`
}

// TypeOutput for type tool
type TypeOutput struct {
	Status string `json:"status"`
}

// ExtractInput for extract tool
type ExtractInput struct {
	Selector string `json:"selector" jsonschema:"required,description=CSS selector of element to extract text from"`
}

// ExtractOutput for extract tool
type ExtractOutput struct {
	Text string `json:"text"`
}

// ScreenshotInput for screenshot tool (empty - no params needed)
type ScreenshotInput struct{}

// ScreenshotOutput for screenshot tool
type ScreenshotOutput struct {
	Screenshot string `json:"screenshot" jsonschema:"description=Base64-encoded PNG screenshot"`
	Format     string `json:"format"`
}

// ExecuteInput for execute tool
type ExecuteInput struct {
	Script string `json:"script" jsonschema:"required,description=JavaScript code to execute on the page"`
}

// ExecuteOutput for execute tool
type ExecuteOutput struct {
	Result string `json:"result"`
}

// HTMLInput for html tool (empty - no params needed)
type HTMLInput struct{}

// HTMLOutput for html tool
type HTMLOutput struct {
	HTML string `json:"html"`
	URL  string `json:"url"`
}

// New creates a new MCP server for browser automation
func New(ctx context.Context) (*Server, error) {
	implementation := &mcp.Implementation{
		Name:    "mime",
		Version: "0.1.0",
	}

	srv := mcp.NewServer(implementation, nil)

	s := &Server{
		server: srv,
		ctx:    ctx,
	}

	// Initialize browser
	br, err := browser.New(ctx, &browser.Options{
		Headless: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create browser: %w", err)
	}
	s.browser = br

	// Register tools
	s.registerTools()

	return s, nil
}

// registerTools registers all MCP tools
func (s *Server) registerTools() {
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
			Description: "Click an element on the page by CSS selector",
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
			data, err := s.browser.Screenshot()
			if err != nil {
				return nil, ScreenshotOutput{}, fmt.Errorf("screenshot failed: %w", err)
			}
			return nil, ScreenshotOutput{Screenshot: data, Format: "base64-png"}, nil
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
			resultJSON, _ := json.Marshal(result)
			return nil, ExecuteOutput{Result: string(resultJSON)}, nil
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
}

// Run starts the MCP server with the given transport
func (s *Server) Run(transport mcp.Transport) error {
	return s.server.Run(s.ctx, transport)
}

// Close closes the browser and cleans up resources
func (s *Server) Close() error {
	if s.browser != nil {
		return s.browser.Close()
	}
	return nil
}
