package mime

import (
	"context"

	"golang.org/x/time/rate"
)

// RateLimiter wraps rate.Limiter
type RateLimiter struct {
	limiter *rate.Limiter
}

// NewRateLimiter creates a new rate limiter
// rps: requests per second
// burst: maximum burst size
func NewRateLimiter(rps float64, burst int) *RateLimiter {
	return &RateLimiter{
		limiter: rate.NewLimiter(rate.Limit(rps), burst),
	}
}

// Wait waits for a token
func (r *RateLimiter) Wait(ctx context.Context) error {
	return r.limiter.Wait(ctx)
}

// DefaultRateLimiter returns a sensible default (5 requests/second, 1 burst)
func DefaultRateLimiter() *RateLimiter {
	return NewRateLimiter(5, 1)
}
