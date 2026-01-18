package mime

import (
	"testing"
	"time"
)

func TestBrowserOptions(t *testing.T) {
	opts := &BrowserOptions{
		Headless: true,
		Timeout:  30 * time.Second,
	}

	if opts.Headless != true {
		t.Error("Headless should be true")
	}

	if opts.Timeout != 30*time.Second {
		t.Error("Timeout should be 30 seconds")
	}
}

func TestWaitForArgs(t *testing.T) {
	args := WaitForArgs{
		Selector: "text=Login",
	}

	if args.Selector != "text=Login" {
		t.Errorf("Expected selector 'text=Login', got '%s'", args.Selector)
	}
}

func TestScrollArgs(t *testing.T) {
	args := ScrollArgs{
		Selector: "",
		X:        0,
		Y:        500,
	}

	if args.X != 0 {
		t.Errorf("Expected X=0, got %d", args.X)
	}

	if args.Y != 500 {
		t.Errorf("Expected Y=500, got %d", args.Y)
	}
}

func TestHoverArgs(t *testing.T) {
	args := HoverArgs{
		Selector: "#button",
	}

	if args.Selector != "#button" {
		t.Errorf("Expected selector '#button', got '%s'", args.Selector)
	}
}
