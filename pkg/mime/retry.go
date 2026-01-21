package mime

import (
	"time"
)

// RetryConfig configures retry behavior
type RetryConfig struct {
	MaxAttempts  int
	InitialDelay time.Duration
	MaxDelay     time.Duration
}

// DefaultRetryConfig returns sensible defaults for retry
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxAttempts:  3,
		InitialDelay: 100 * time.Millisecond,
		MaxDelay:     2 * time.Second,
	}
}

// WithRetry executes a function with retry logic
func WithRetry[T any](cfg *RetryConfig, fn func() (T, error)) (T, error) {
	var result T
	var err error
	delay := cfg.InitialDelay

	for i := 0; i < cfg.MaxAttempts; i++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
		if !isRetryable(err) {
			return result, err
		}
		if i < cfg.MaxAttempts-1 {
			time.Sleep(delay)
			delay = min(delay*2, cfg.MaxDelay)
		}
	}
	return result, err
}

// WithRetryVoid executes a void function with retry logic
func WithRetryVoid(cfg *RetryConfig, fn func() error) error {
	_, err := WithRetry(cfg, func() (struct{}, error) {
		return struct{}{}, fn()
	})
	return err
}

// isRetryable determines if an error should be retried
func isRetryable(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()

	// Retry on common transient errors
	retryablePatterns := []string{
		"timeout",
		"connection refused",
		"connection reset",
		"temporary failure",
		"context deadline",
	}

	for _, pattern := range retryablePatterns {
		if contains(errStr, pattern) {
			return true
		}
	}

	return false
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
