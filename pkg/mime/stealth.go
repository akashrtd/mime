package mime

import (
	"fmt"
)

// StealthOptions configures anti-detection measures
type StealthOptions struct {
	UserAgent         string `json:"user_agent"`
	HideWebDriver     bool   `json:"hide_webdriver"`
	DisableAutomation bool   `json:"disable_automation"` // Disable automation-controlled flag
}

// DefaultStealthOptions returns common stealth settings
func DefaultStealthOptions() *StealthOptions {
	return &StealthOptions{
		UserAgent:     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		HideWebDriver: true,
	}
}

// ApplyStealth applies stealth settings to the browser
func (b *Browser) ApplyStealth(opts *StealthOptions) error {
	if opts == nil {
		opts = DefaultStealthOptions()
	}

	// 1. Override User-Agent
	if opts.UserAgent != "" {
		// Use rod's SetUserAgent? Or just via CDP
		// browser.go doesn't expose easy user agent setter yet, we can do it via page.SetUserAgent
		// But b.page isn't locked.

		// For now, let's assume we do this on the current page
		// NOTE: Ideally this should be set at Browser context level on creation
	}

	// 2. Hide webdriver flag
	if opts.HideWebDriver {
		script := `() => {
			Object.defineProperty(navigator, 'webdriver', {
				get: () => false,
			});
		}`
		if _, err := b.Execute(script); err != nil {
			return fmt.Errorf("failed to hide webdriver: %w", err)
		}
	}

	// 3. More stealth techniques can be added here
	// e.g. masking plugins, languages, etc.

	return nil
}
