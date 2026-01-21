package mime

import (
	"context"
	"fmt"
	"sync"
)

// BrowserPool manages a pool of browser instances
type BrowserPool struct {
	browsers []*Browser
	mu       sync.Mutex
	maxSize  int
	opts     *BrowserOptions
}

// NewBrowserPool creates a new browser pool
func NewBrowserPool(ctx context.Context, size int, opts *BrowserOptions) (*BrowserPool, error) {
	if size <= 0 {
		size = 1
	}

	pool := &BrowserPool{
		browsers: make([]*Browser, 0, size),
		maxSize:  size,
		opts:     opts,
	}

	// Pre-warm with one browser
	if size > 0 {
		br, err := NewBrowser(ctx, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to pre-warm pool: %w", err)
		}
		pool.browsers = append(pool.browsers, br)
	}

	return pool, nil
}

// Get retrieves a browser from the pool or creates a new one
func (p *BrowserPool) Get(ctx context.Context) (*Browser, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.browsers) > 0 {
		// Pop the last browser
		n := len(p.browsers)
		br := p.browsers[n-1]
		p.browsers = p.browsers[:n-1]

		// Check if browser is still alive?
		// For now assume yes, or we could add a check
		return br, nil
	}

	// Pool empty, create new
	return NewBrowser(ctx, p.opts)
}

// Put returns a browser to the pool
func (p *BrowserPool) Put(br *Browser) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.browsers) < p.maxSize {
		p.browsers = append(p.browsers, br)
	} else {
		// Pool full, close the browser
		go br.Close()
	}
}

// Close closes all browsers in the pool
func (p *BrowserPool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	var err error
	for _, br := range p.browsers {
		if e := br.Close(); e != nil {
			err = e // Return last error
		}
	}
	p.browsers = nil
	return err
}
