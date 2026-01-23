package mime

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Tracer records browser actions
type Tracer struct {
	Steps []TraceStep `json:"steps"`
	mu    sync.Mutex
	Dir   string
}

// TraceStep represents a single action step
type TraceStep struct {
	Timestamp  time.Time `json:"timestamp"`
	Action     string    `json:"action"`
	Selector   string    `json:"selector,omitempty"`
	URL        string    `json:"url,omitempty"`
	Screenshot string    `json:"screenshot,omitempty"` // Path to screenshot file
}

// NewTracer creates a new tracer
func NewTracer(dir string) (*Tracer, error) {
	if dir == "" {
		dir = "traces"
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create trace directory: %w", err)
	}
	return &Tracer{
		Dir:   dir,
		Steps: make([]TraceStep, 0),
	}, nil
}

// Capture captures a step
func (t *Tracer) Capture(b *Browser, action, selector string) {
	if t == nil {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()

	step := TraceStep{
		Timestamp: time.Now(),
		Action:    action,
		Selector:  selector,
		URL:       b.URL(),
	}

	// Capture screenshot
	// Use a simplified version to avoid infinite recursion if screenshot triggers trace (it shouldn't)
	// We call internal method if possible, or just public one
	// NOTE: We should modify Browser to expose raw *rod.Page or have an internal screenshot method to avoid loops?
	// Actually public Screenshot() is fine as long as we don't instrument Screenshot() itself.
	data, err := b.Screenshot()
	if err == nil {
		// Save screenshot to file
		filename := fmt.Sprintf("step_%d_%s.png", len(t.Steps), time.Unix(step.Timestamp.Unix(), 0).Format("150405"))
		path := filepath.Join(t.Dir, filename)
		// data is base64, decode it
		// wait, b.Screenshot returns base64 string. We should decode it to save space?
		// Or just save base64? Traces usually save raw images.
		// Let's rely on b.page.Screenshot(false, nil) directly if we could to separate concerns,
		// but b.Screenshot() encapsulates error handling nicely.
		// Let's stick with decoding the base64 string or better yet, make b.Screenshot() helpful.
		// For now simple write.
		if err := os.WriteFile(path+".b64", []byte(data), 0644); err == nil {
			step.Screenshot = filename + ".b64"
		}
	}

	t.Steps = append(t.Steps, step)

	// Auto-save index
	t.saveIndex()
}

func (t *Tracer) saveIndex() {
	path := filepath.Join(t.Dir, "trace.json")
	data, _ := json.MarshalIndent(t, "", "  ")
	_ = os.WriteFile(path, data, 0644)
}
