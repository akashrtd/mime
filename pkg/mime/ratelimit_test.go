package mime

import (
	"context"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	// Configure: 10 requests per second (100ms interval), burst 1
	limiter := NewRateLimiter(10, 1)

	ctx := context.Background()
	start := time.Now()

	// First request should happen immediately
	if err := limiter.Wait(ctx); err != nil {
		t.Fatalf("First wait failed: %v", err)
	}

	// Second request should wait ~100ms
	if err := limiter.Wait(ctx); err != nil {
		t.Fatalf("Second wait failed: %v", err)
	}

	elapsed := time.Since(start)
	if elapsed < 80*time.Millisecond {
		t.Errorf("Rate limiter didn't wait enough. Elapsed: %v", elapsed)
	}
}
