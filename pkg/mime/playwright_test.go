package mime

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestPlaywrightFeatures(t *testing.T) {
	// Setup temporary trace dir
	traceDir := "test_traces"
	os.RemoveAll(traceDir)
	defer os.RemoveAll(traceDir)

	ctx := context.Background()
	opts := &BrowserOptions{
		Headless: true,
		TraceDir: traceDir,
	}

	b, err := NewBrowser(ctx, opts)
	if err != nil {
		t.Fatalf("Failed to create browser: %v", err)
	}
	defer b.Close()

	// 1. Test Network Interception (Block Resources)
	if err := b.BlockResources(); err != nil {
		t.Fatalf("Failed to block resources: %v", err)
	}

	// 2. Test Navigation (Tracer should capture)
	if err := b.Navigate("https://example.com"); err != nil {
		t.Fatalf("Failed to navigate: %v", err)
	}

	// 3. Test Auto-Wait (Implicitly tested by Click)
	// We can't easily simulate unstable element without a custom HTML page,
	// but we can check if Click returns error on non-existent
	// err = b.Click("#non-existent") -- this would timeout

	// 4. Verify Traces
	time.Sleep(500 * time.Millisecond) // Allow async trace run? No, it's inside Click/Navigate but screenshot might be slow

	// Check if directory exists
	if _, err := os.Stat(traceDir); os.IsNotExist(err) {
		t.Fatalf("Trace directory not created")
	}

	// Check for trace.json
	if _, err := os.Stat(traceDir + "/trace.json"); os.IsNotExist(err) {
		t.Fatalf("trace.json not created")
	}
}
