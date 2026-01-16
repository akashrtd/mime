package browser

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// Browser represents a browser automation instance
type Browser struct {
	browser *rod.Browser
	page    *rod.Page
	ctx     context.Context
}

// Options configures browser behavior
type Options struct {
	Headless bool
	Timeout  time.Duration
}

// New creates a new Browser instance
func New(ctx context.Context, opts *Options) (*Browser, error) {
	if opts == nil {
		opts = &Options{
			Headless: true,
			Timeout:  30 * time.Second,
		}
	}

	// Launch browser
	l := launcher.New().Headless(opts.Headless)
	url := l.MustLaunch()

	browser := rod.New().ControlURL(url).Context(ctx)
	if err := browser.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to browser: %w", err)
	}

	// Create initial page
	page, err := browser.Page(proto.TargetCreateTarget{})
	if err != nil {
		return nil, fmt.Errorf("failed to create page: %w", err)
	}

	return &Browser{
		browser: browser,
		page:    page,
		ctx:     ctx,
	}, nil
}

// Navigate navigates to a URL
func (b *Browser) Navigate(url string) error {
	return b.page.Navigate(url)
}

// Click clicks an element by selector
func (b *Browser) Click(selector string) error {
	elem, err := b.page.Element(selector)
	if err != nil {
		return fmt.Errorf("element not found: %s: %w", selector, err)
	}
	return elem.Click(proto.InputMouseButtonLeft, 1)
}

// Type types text into an element
func (b *Browser) Type(selector, text string) error {
	elem, err := b.page.Element(selector)
	if err != nil {
		return fmt.Errorf("element not found: %s: %w", selector, err)
	}
	return elem.Input(text)
}

// Extract extracts text from an element
func (b *Browser) Extract(selector string) (string, error) {
	elem, err := b.page.Element(selector)
	if err != nil {
		return "", fmt.Errorf("element not found: %s: %w", selector, err)
	}
	text, err := elem.Text()
	if err != nil {
		return "", fmt.Errorf("failed to extract text: %w", err)
	}
	return text, nil
}

// ExtractAttr extracts an attribute from an element
func (b *Browser) ExtractAttr(selector, attr string) (string, error) {
	elem, err := b.page.Element(selector)
	if err != nil {
		return "", fmt.Errorf("element not found: %s: %w", selector, err)
	}
	val, err := elem.Attribute(attr)
	if err != nil {
		return "", fmt.Errorf("failed to extract attribute: %w", err)
	}
	if val == nil {
		return "", nil
	}
	return *val, nil
}

// Screenshot captures a screenshot and returns base64-encoded PNG
func (b *Browser) Screenshot() (string, error) {
	data, err := b.page.Screenshot(false, nil)
	if err != nil {
		return "", fmt.Errorf("failed to capture screenshot: %w", err)
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// WaitFor waits for an element to appear
func (b *Browser) WaitFor(selector string) error {
	_, err := b.page.Element(selector)
	return err
}

// Execute runs JavaScript code on the page
func (b *Browser) Execute(script string) (interface{}, error) {
	result, err := b.page.Eval(script)
	if err != nil {
		return nil, fmt.Errorf("failed to execute script: %w", err)
	}
	return result.Value, nil
}

// HTML returns the current page HTML
func (b *Browser) HTML() (string, error) {
	html, err := b.page.HTML()
	if err != nil {
		return "", fmt.Errorf("failed to get HTML: %w", err)
	}
	return html, nil
}

// URL returns the current page URL
func (b *Browser) URL() string {
	info := b.page.MustInfo()
	return info.URL
}

// Title returns the current page title
func (b *Browser) Title() (string, error) {
	info, err := b.page.Info()
	if err != nil {
		return "", fmt.Errorf("failed to get page info: %w", err)
	}
	return info.Title, nil
}

// Close closes the browser
func (b *Browser) Close() error {
	return b.browser.Close()
}
