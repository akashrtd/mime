package mime

import (
	"context"
	"testing"
)

func TestBrowserPool(t *testing.T) {
	ctx := context.Background()
	opts := &BrowserOptions{
		Headless: true,
	}

	t.Run("Initialize and Pre-warm", func(t *testing.T) {
		pool, err := NewBrowserPool(ctx, 2, opts)
		if err != nil {
			t.Fatalf("Failed to create pool: %v", err)
		}
		defer pool.Close()

		if len(pool.browsers) != 1 {
			t.Errorf("Expected 1 pre-warmed browser, got %d", len(pool.browsers))
		}
	})

	t.Run("Get and Put", func(t *testing.T) {
		pool, err := NewBrowserPool(ctx, 2, opts) // Size 2
		if err != nil {
			t.Fatalf("Failed to create pool: %v", err)
		}
		defer pool.Close()

		// Get 1 (pre-warmed)
		b1, err := pool.Get(ctx)
		if err != nil {
			t.Fatalf("Get 1 failed: %v", err)
		}
		if len(pool.browsers) != 0 {
			t.Errorf("Pool should be empty after getting pre-warmed browser")
		}

		// Get 2 (new)
		b2, err := pool.Get(ctx)
		if err != nil {
			t.Fatalf("Get 2 failed: %v", err)
		}

		// Put both back
		pool.Put(b1)
		pool.Put(b2)

		if len(pool.browsers) != 2 {
			t.Errorf("Pool size should be 2, got %d", len(pool.browsers))
		}

		// Put overflow (should be closed, pool stays at max size)
		b3, err := NewBrowser(ctx, opts)
		if err != nil {
			t.Fatalf("Failed to create b3: %v", err)
		}
		pool.Put(b3) // Should close b3

		if len(pool.browsers) != 2 {
			t.Errorf("Pool size should stay at max 2, got %d", len(pool.browsers))
		}

		// Ensure b3 is closed? Hard to check without mock.
		// We assume Put closes it.
	})

	t.Run("Close", func(t *testing.T) {
		pool, err := NewBrowserPool(ctx, 1, opts)
		if err != nil {
			t.Fatalf("Failed to create pool: %v", err)
		}

		err = pool.Close()
		if err != nil {
			t.Errorf("Close failed: %v", err)
		}

		if pool.browsers != nil && len(pool.browsers) > 0 {
			t.Error("Pool browsers should be nil or empty after close")
		}
	})
}
