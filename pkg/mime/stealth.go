package mime

import (
	"fmt"

	"github.com/go-rod/rod/lib/proto"
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
		if err := b.page.SetUserAgent(&proto.NetworkSetUserAgentOverride{
			UserAgent: opts.UserAgent,
		}); err != nil {
			return fmt.Errorf("failed to set user agent: %w", err)
		}
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
