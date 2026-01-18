package mime

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

// BrowserOptions configures browser behavior
type BrowserOptions struct {
	Headless bool
	Timeout  time.Duration
}

// NewBrowser creates a new Browser instance
func NewBrowser(ctx context.Context, opts *BrowserOptions) (*Browser, error) {
	if opts == nil {
		opts = &BrowserOptions{
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

// getElement resolves a selector to an element, supporting "text=" prefix
func (b *Browser) getElement(selector string) (*rod.Element, error) {
	// Check for text selector prefix
	if len(selector) > 5 && selector[:5] == "text=" {
		text := selector[5:]
		// Use rod's text search
		elem, err := b.page.ElementR("body", text)
		if err != nil {
			return nil, fmt.Errorf("element with text %q not found: %w", text, err)
		}
		return elem, nil
	}

	// Default to CSS selector
	elem, err := b.page.Element(selector)
	if err != nil {
		return nil, fmt.Errorf("element %q not found: %w", selector, err)
	}
	return elem, nil
}

// Navigate navigates to a URL
func (b *Browser) Navigate(url string) error {
	return b.page.Navigate(url)
}

// Click clicks an element by selector
func (b *Browser) Click(selector string) error {
	elem, err := b.getElement(selector)
	if err != nil {
		return err
	}
	// Ensure element is visible before clicking
	if err := elem.WaitVisible(); err != nil {
		return fmt.Errorf("element not visible: %w", err)
	}
	return elem.Click(proto.InputMouseButtonLeft, 1)
}

// Type types text into an element
func (b *Browser) Type(selector, text string) error {
	elem, err := b.getElement(selector)
	if err != nil {
		return err
	}
	// Select all text first to ensure we replace "input" content if needed,
	// or matches standard "fill" behavior.
	// Rod's Input appends. SelectAll + Input replaces.
	if err := elem.SelectAllText(); err == nil {
		if err := elem.Input(text); err != nil {
			return err
		}
	} else {
		// If SelectAll fails (e.g. not an input), try just Input
		if err := elem.Input(text); err != nil {
			return err
		}
	}
	return nil
}

// Extract extracts text from an element
func (b *Browser) Extract(selector string) (string, error) {
	elem, err := b.getElement(selector)
	if err != nil {
		return "", err
	}
	text, err := elem.Text()
	if err != nil {
		return "", fmt.Errorf("failed to extract text: %w", err)
	}
	return text, nil
}

// ExtractAttr extracts an attribute from an element
func (b *Browser) ExtractAttr(selector, attr string) (string, error) {
	elem, err := b.getElement(selector)
	if err != nil {
		return "", err
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
	// getElement handles both CSS and text selectors, and Rod waits by default
	_, err := b.getElement(selector)
	return err
}

// Hover hovers over an element
func (b *Browser) Hover(selector string) error {
	elem, err := b.getElement(selector)
	if err != nil {
		return err
	}
	return elem.Hover()
}

// Scroll scrolls the window or an element
func (b *Browser) Scroll(selector string, x, y int) error {
	if selector != "" {
		elem, err := b.getElement(selector)
		if err != nil {
			return err
		}
		return elem.ScrollIntoView()
	}
	// Scroll window by x, y
	b.page.Mouse.Scroll(float64(x), float64(y), 0)
	return nil
}

// execute script (unchanged)
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
