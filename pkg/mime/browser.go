package mime

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
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
	logger  *slog.Logger
	limiter *RateLimiter
	tracer  *Tracer
}

// BrowserOptions configures browser behavior
type BrowserOptions struct {
	Headless    bool
	Timeout     time.Duration
	UserDataDir string          // Path to persist session data (cookies, localStorage)
	Stealth     *StealthOptions // Anti-detection configuration
	Logger      *slog.Logger    // Structured logger
	RateLimit   *RateLimiter    // Optional rate limiter
	TraceDir    string          // Directory for traces (if empty, tracing disabled)
}

// PageMetadata contains page meta information
type PageMetadata struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	URL         string            `json:"url"`
	Canonical   string            `json:"canonical,omitempty"`
	OG          map[string]string `json:"og,omitempty"`
}

// Link represents a hyperlink on the page
type Link struct {
	URL  string `json:"url"`
	Text string `json:"text"`
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
	// Launch browser with optimized flags for speed
	l := launcher.New().
		Headless(opts.Headless).
		Set("disable-gpu", "true").
		Set("no-sandbox", "true").
		Set("disable-setuid-sandbox", "true").
		Set("disable-dev-shm-usage", "true").
		Set("disable-infobars", "true").
		Set("no-first-run", "true").
		Set("no-default-browser-check", "true")

	if opts.UserDataDir != "" {
		l = l.UserDataDir(opts.UserDataDir)
	}
	url := l.MustLaunch()

	// Setup logger
	logger := opts.Logger
	if logger == nil {
		logger = slog.Default()
	}

	logger.Debug("launching browser", "headless", opts.Headless, "userDataDir", opts.UserDataDir)

	browser := rod.New().ControlURL(url).Context(ctx)
	if err := browser.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to browser: %w", err)
	}

	// Create initial page
	page, err := browser.Page(proto.TargetCreateTarget{})
	if err != nil {
		return nil, fmt.Errorf("failed to create page: %w", err)
	}

	b := &Browser{
		browser: browser,
		page:    page,
		ctx:     ctx,
		logger:  logger,
		limiter: opts.RateLimit,
	}

	if opts.TraceDir != "" {
		tracer, err := NewTracer(opts.TraceDir)
		if err != nil {
			return nil, fmt.Errorf("failed to init tracer: %w", err)
		}
		b.tracer = tracer
		b.logger.Info("tracing enabled", "dir", opts.TraceDir)
	}

	if opts.Stealth != nil {
		if err := b.ApplyStealth(opts.Stealth); err != nil {
			return nil, fmt.Errorf("failed to apply stealth settings: %w", err)
		}
	}

	return b, nil
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
	// Use ElementR for text, but for standard selectors we want to support Shadow DOM piercing if possible.
	// Rod's Element() is querySelector.
	// To better support Playwright-style >> or just finding elements in shadow roots,
	// we can try a recursive search or just stick to Rod's defaults but ensure we wait.
	// For this iteration, we keep it simple but acknowledge Shadow DOM needs might require js search.
	// Actually Rod supports piercing shadow dom with `Element` if the selector is crafted right?
	// No, usually you need `ShadowRoot()`.
	// Let's rely on Rod's retry logic which is quite good.
	// IMPORTANT: Playwright-style deep selector is not default in Rod.

	elem, err := b.page.Element(selector)
	if err != nil {
		return nil, fmt.Errorf("element %q not found: %w", selector, err)
	}
	return elem, nil
}

// Navigate navigates to a URL
func (b *Browser) Navigate(url string) error {
	if b.limiter != nil {
		if err := b.limiter.Wait(b.ctx); err != nil {
			return fmt.Errorf("rate limit wait failed: %w", err)
		}
	}

	// Security: Block file:// scheme
	if strings.HasPrefix(url, "file://") {
		return fmt.Errorf("security error: file:// scheme is not allowed")
	}

	b.logger.Info("navigating", "url", url)
	if b.tracer != nil {
		b.tracer.Capture(b, "navigate", url)
	}
	return b.page.Navigate(url)
}

// Click clicks an element by selector
func (b *Browser) Click(selector string) error {
	if b.tracer != nil {
		defer b.tracer.Capture(b, "click", selector)
	}
	elem, err := b.getElement(selector)
	if err != nil {
		return err
	}
	// Ensure element is visible and stable
	if err := elem.WaitVisible(); err != nil {
		return fmt.Errorf("element not visible: %w", err)
	}
	if err := elem.WaitStable(100 * time.Millisecond); err != nil {
		b.logger.Warn("element not stable, clicking anyway", "selector", selector, "error", err)
	}
	if err := elem.WaitEnabled(); err != nil {
		return fmt.Errorf("element not enabled: %w", err)
	}
	return elem.Click(proto.InputMouseButtonLeft, 1)
}

// Type types text into an element
func (b *Browser) Type(selector, text string) error {
	if b.tracer != nil {
		defer b.tracer.Capture(b, "type", selector)
	}
	elem, err := b.getElement(selector)
	if err != nil {
		return err
	}
	// Ensure element is visible and stable
	if err := elem.WaitVisible(); err != nil {
		return fmt.Errorf("element not visible: %w", err)
	}
	if err := elem.WaitStable(100 * time.Millisecond); err != nil {
		b.logger.Warn("element not stable, typing anyway", "selector", selector, "error", err)
	}
	if err := elem.WaitEnabled(); err != nil {
		return fmt.Errorf("element not enabled: %w", err)
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
// Screenshot captures a screenshot and returns base64-encoded string
func (b *Browser) Screenshot(format string, quality int) (string, error) {
	req := &proto.PageCaptureScreenshot{
		Format: proto.PageCaptureScreenshotFormatPng,
	}

	if format == "jpeg" || format == "jpg" {
		req.Format = proto.PageCaptureScreenshotFormatJpeg
		if quality > 0 && quality <= 100 {
			q := quality
			req.Quality = &q
		} else {
			q := 80 // Default JPEG quality for speed
			req.Quality = &q
		}
	}

	data, err := req.Call(b.page)
	if err != nil {
		return "", fmt.Errorf("failed to capture screenshot: %w", err)
	}
	return base64.StdEncoding.EncodeToString(data.Data), nil
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
	if b.limiter != nil {
		if err := b.limiter.Wait(b.ctx); err != nil {
			return nil, fmt.Errorf("rate limit wait failed: %w", err)
		}
	}
	// Wrap script in a function to work with rod's Eval
	// rod's Eval expects: `function() { ... }` format
	wrappedScript := fmt.Sprintf("() => { return %s; }", script)
	result, err := b.page.Eval(wrappedScript)
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

// GetCookies returns all cookies for the current page
func (b *Browser) GetCookies() ([]*proto.NetworkCookie, error) {
	cookies, err := b.page.Cookies(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get cookies: %w", err)
	}
	return cookies, nil
}

// SetCookies sets cookies for the browser
func (b *Browser) SetCookies(cookies []*proto.NetworkCookieParam) error {
	return b.page.SetCookies(cookies)
}

// ClearCookies clears all cookies
func (b *Browser) ClearCookies() error {
	return b.page.SetCookies(nil)
}

// Metadata extracts page metadata (title, description, og tags)
func (b *Browser) Metadata() (*PageMetadata, error) {
	script := `() => {
		const og = {};
		document.querySelectorAll('meta[property^="og:"]').forEach(m => {
			const key = m.getAttribute('property').replace('og:', '');
			og[key] = m.content;
		});
		return {
			title: document.title || '',
			description: document.querySelector('meta[name="description"]')?.content || '',
			canonical: document.querySelector('link[rel="canonical"]')?.href || '',
			og: og
		};
	}`

	result, err := b.page.Eval(script)
	if err != nil {
		return nil, fmt.Errorf("failed to extract metadata: %w", err)
	}

	meta := &PageMetadata{
		URL: b.URL(),
	}

	// Use JSON for type-safe conversion
	jsonBytes, err := result.Value.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal metadata: %w", err)
	}

	var data struct {
		Title       string            `json:"title"`
		Description string            `json:"description"`
		Canonical   string            `json:"canonical"`
		OG          map[string]string `json:"og"`
	}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, fmt.Errorf("failed to parse metadata: %w", err)
	}

	meta.Title = data.Title
	meta.Description = data.Description
	meta.Canonical = data.Canonical
	meta.OG = data.OG

	return meta, nil
}

// Links extracts all links from the current page
func (b *Browser) Links() ([]Link, error) {
	script := `() => {
		return [...document.querySelectorAll('a[href]')].map(a => ({
			url: a.href,
			text: a.textContent.trim().slice(0, 200)
		})).filter(l => l.url.startsWith('http'));
	}`

	result, err := b.page.Eval(script)
	if err != nil {
		return nil, fmt.Errorf("failed to extract links: %w", err)
	}

	// Use JSON for type-safe conversion
	jsonBytes, err := result.Value.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal links: %w", err)
	}

	var links []Link
	if err := json.Unmarshal(jsonBytes, &links); err != nil {
		return nil, fmt.Errorf("failed to parse links: %w", err)
	}

	return links, nil
}
