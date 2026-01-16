package mcp

import (
	"context"
	"encoding/base64"
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

// New creates a new MCP server
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

	// Register resources
	s.registerResources()

	// Register tools
	s.registerTools()

	return s, nil
}

// registerResources registers MCP resources
func (s *Server) registerResources() {
	// Page HTML resource
	mcp.AddResource(s.server,
		&mcp.Resource{
			URI:         "browser://page",
			Name:        "Current Page HTML",
			Description: "The HTML content of the current browser page",
			MIMEType:    "text/html",
		},
		func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
			html, err := s.browser.HTML()
			if err != nil {
				return nil, err
			}
			return &mcp.ReadResourceResult{
				Contents: []mcp.ResourceContents{
					&mcp.TextResourceContents{
						ResourceContents: mcp.ResourceContents{
							URI:      "browser://page",
							MIMEType: "text/html",
						},
						Text: html,
					},
				},
			}, nil
		},
	)

	// Screenshot resource
	mcp.AddResource(s.server,
		&mcp.Resource{
			URI:         "browser://screenshot",
			Name:        "Current Page Screenshot",
			Description: "A screenshot of the current browser page (base64 PNG)",
			MIMEType:    "image/png",
		},
		func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
			screenshot, err := s.browser.Screenshot()
			if err != nil {
				return nil, err
			}
			// Decode base64 to bytes
			data, err := base64.StdEncoding.DecodeString(screenshot)
			if err != nil {
				return nil, err
			}
			return &mcp.ReadResourceResult{
				Contents: []mcp.ResourceContents{
					&mcp.BlobResourceContents{
						ResourceContents: mcp.ResourceContents{
							URI:      "browser://screenshot",
							MIMEType: "image/png",
						},
						Blob: data,
					},
				},
			}, nil
		},
	)
}

// NavigateInput defines the input for navigate tool
type NavigateInput struct {
	URL string `json:"url" jsonschema:"required,description=URL to navigate to"`
}

// NavigateOutput defines the output for navigate tool
type NavigateOutput struct {
	Status string `json:"status"`
	URL    string `json:"url"`
}

// ClickInput defines the input for click tool
type ClickInput struct {
	Selector string `json:"selector" jsonschema:"required,description=CSS selector of element to click"`
}

// ClickOutput defines the output for click tool
type ClickOutput struct {
	Status   string `json:"status"`
	Selector string `json:"selector"`
}

// TypeInput defines the input for type tool
type TypeInput struct {
	Selector string `json:"selector" jsonschema:"required,description=CSS selector of element to type into"`
	Text     string `json:"text" jsonschema:"required,description=Text to type"`
}

// TypeOutput defines the output for type tool
type TypeOutput struct {
	Status string `json:"status"`
}

// ExtractInput defines the input for extract tool
type ExtractInput struct {
	Selector string `json:"selector" jsonschema:"required,description=CSS selector of element to extract from"`
}

// ExtractOutput defines the output for extract tool
type ExtractOutput struct {
	Text string `json:"text"`
}

// ScreenshotOutput defines the output for screenshot tool
type ScreenshotOutput struct {
	Screenshot string `json:"screenshot" jsonschema:"description=Base64-encoded PNG screenshot"`
	Format     string `json:"format"`
}

// ExecuteInput defines the input for execute tool
type ExecuteInput struct {
	Script string `json:"script" jsonschema:"required,description=JavaScript code to execute"`
}

// ExecuteOutput defines the output for execute tool
type ExecuteOutput struct {
	Result string `json:"result"`
}

// registerTools registers MCP tools
func (s *Server) registerTools() {
	// Navigate tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "navigate",
			Description: "Navigate to a URL in the browser",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input NavigateInput) (*mcp.CallToolResult, NavigateOutput, error) {
			if err := s.browser.Navigate(input.URL); err != nil {
				return nil, NavigateOutput{}, err
			}
			return nil, NavigateOutput{Status: "success", URL: input.URL}, nil
		},
	)

	// Click tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "click",
			Description: "Click an element by CSS selector",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input ClickInput) (*mcp.CallToolResult, ClickOutput, error) {
			if err := s.browser.Click(input.Selector); err != nil {
				return nil, ClickOutput{}, err
			}
			return nil, ClickOutput{Status: "success", Selector: input.Selector}, nil
		},
	)

	// Type tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "type",
			Description: "Type text into an element",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input TypeInput) (*mcp.CallToolResult, TypeOutput, error) {
			if err := s.browser.Type(input.Selector, input.Text); err != nil {
				return nil, TypeOutput{}, err
			}
			return nil, TypeOutput{Status: "success"}, nil
		},
	)

	// Extract tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "extract",
			Description: "Extract text from an element",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input ExtractInput) (*mcp.CallToolResult, ExtractOutput, error) {
			text, err := s.browser.Extract(input.Selector)
			if err != nil {
				return nil, ExtractOutput{}, err
			}
			return nil, ExtractOutput{Text: text}, nil
		},
	)

	// Screenshot tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "screenshot",
			Description: "Capture a screenshot of the current page",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, ScreenshotOutput, error) {
			data, err := s.browser.Screenshot()
			if err != nil {
				return nil, ScreenshotOutput{}, err
			}
			return nil, ScreenshotOutput{Screenshot: data, Format: "base64-png"}, nil
		},
	)

	// Execute tool
	mcp.AddTool(s.server,
		&mcp.Tool{
			Name:        "execute",
			Description: "Execute JavaScript code on the page",
		},
		func(ctx context.Context, req *mcp.CallToolRequest, input ExecuteInput) (*mcp.CallToolResult, ExecuteOutput, error) {
			result, err := s.browser.Execute(input.Script)
			if err != nil {
				return nil, ExecuteOutput{}, err
			}
			resultJSON, _ := json.Marshal(result)
			return nil, ExecuteOutput{Result: string(resultJSON)}, nil
		},
	)
}

// Run starts the MCP server
func (s *Server) Run(transport mcp.Transport) error {
	return s.server.Run(s.ctx, transport)
}

// Close closes the server and browser
func (s *Server) Close() error {
	if s.browser != nil {
		return s.browser.Close()
	}
	return nil
}
