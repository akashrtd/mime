package mime

import (
	"context"
	"time"
)

// MIME provides a fluent API for browser automation
type MIME struct {
	browser *Browser
	ctx     context.Context
}

// New creates a new MIME instance
func New(ctx context.Context) (*MIME, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	br, err := NewBrowser(ctx, &BrowserOptions{
		Headless: true,
		Timeout:  30 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return &MIME{
		browser: br,
		ctx:     ctx,
	}, nil
}

// NewWithOptions creates a new MIME instance with custom options
func NewWithOptions(ctx context.Context, opts *BrowserOptions) (*MIME, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	br, err := NewBrowser(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &MIME{
		browser: br,
		ctx:     ctx,
	}, nil
}

// Navigate navigates to a URL
func (m *MIME) Navigate(url string) error {
	return m.browser.Navigate(url)
}

// Click clicks an element
func (m *MIME) Click(selector string) error {
	return m.browser.Click(selector)
}

// Type types text into an element
func (m *MIME) Type(selector, text string) error {
	return m.browser.Type(selector, text)
}

// Extract extracts text from an element
func (m *MIME) Extract(selector string) (string, error) {
	return m.browser.Extract(selector)
}

// ExtractAttr extracts an attribute from an element
func (m *MIME) ExtractAttr(selector, attr string) (string, error) {
	return m.browser.ExtractAttr(selector, attr)
}

// Screenshot captures a screenshot
func (m *MIME) Screenshot() (string, error) {
	return m.browser.Screenshot("", 0)
}

// WaitFor waits for an element to appear
func (m *MIME) WaitFor(selector string) error {
	return m.browser.WaitFor(selector)
}

// Execute runs JavaScript code
func (m *MIME) Execute(script string) (interface{}, error) {
	return m.browser.Execute(script)
}

// HTML returns the current page HTML
func (m *MIME) HTML() (string, error) {
	return m.browser.HTML()
}

// URL returns the current page URL
func (m *MIME) URL() string {
	return m.browser.URL()
}

// Title returns the current page title
func (m *MIME) Title() (string, error) {
	return m.browser.Title()
}

// Close closes the browser
func (m *MIME) Close() error {
	return m.browser.Close()
}
