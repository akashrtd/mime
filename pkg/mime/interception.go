package mime

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// InterceptOptions configures network interception
type InterceptOptions struct {
	BlockImages bool     `json:"block_images"`
	BlockFonts  bool     `json:"block_fonts"`
	BlockMedia  bool     `json:"block_media"`
	BlockURL    []string `json:"block_url,omitempty"` // Block specific URL patterns
}

// EnableInterception enables network interception on the browser
func (b *Browser) EnableInterception(opts *InterceptOptions) error {
	router := b.browser.HijackRequests()

	// Default blockers if options are minimal
	if opts.BlockImages {
		router.MustAdd("*.png", httpResponseBlock)
		router.MustAdd("*.jpg", httpResponseBlock)
		router.MustAdd("*.jpeg", httpResponseBlock)
		router.MustAdd("*.gif", httpResponseBlock)
		router.MustAdd("*.webp", httpResponseBlock)
		router.MustAdd("*.svg", httpResponseBlock)
	}

	if opts.BlockFonts {
		router.MustAdd("*.woff", httpResponseBlock)
		router.MustAdd("*.woff2", httpResponseBlock)
		router.MustAdd("*.ttf", httpResponseBlock)
		router.MustAdd("*.otf", httpResponseBlock)
	}

	if opts.BlockMedia {
		router.MustAdd("*.mp4", httpResponseBlock)
		router.MustAdd("*.mp3", httpResponseBlock)
		router.MustAdd("*.wav", httpResponseBlock)
		router.MustAdd("*.webm", httpResponseBlock)
	}

	// Custom URL blockers
	for _, pattern := range opts.BlockURL {
		router.MustAdd(pattern, httpResponseBlock)
	}

	// Start hijacking
	go router.Run()
	return nil
}

// BlockResources is a shorthand for common blocking
func (b *Browser) BlockResources() error {
	return b.EnableInterception(&InterceptOptions{
		BlockImages: true,
		BlockFonts:  true,
		BlockMedia:  true,
	})
}

// httpResponseBlock simulates a failed request
func httpResponseBlock(ctx *rod.Hijack) {
	ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
}

// LogRequests logs all network requests (useful for debugging/tracing)
func (b *Browser) LogRequests() {
	go b.browser.EachEvent(func(e *proto.NetworkRequestWillBeSent) {
		b.logger.Debug("networking", "url", e.Request.URL, "method", e.Request.Method)
	})
}
