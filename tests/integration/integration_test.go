package tests

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/akashrtd/mime/pkg/mime"
)

func TestIntegration_ObserveAndAct(t *testing.T) {
	// 1. Start Test Server
	server := StartTestServer(t)
	targetURL := server.URL + "/forms.html"

	// 2. Launch Browser
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	opts := &mime.BrowserOptions{
		Headless: true,
	}
	browser, err := mime.NewBrowser(ctx, opts)
	if err != nil {
		t.Fatalf("Failed to create browser: %v", err)
	}
	defer browser.Close()

	// 3. Navigate
	if err := browser.Navigate(targetURL); err != nil {
		t.Fatalf("Failed to navigate: %v", err)
	}

	// 3b. Title() must reflect the loaded page immediately after Navigate,
	// not a stale "about:blank" from before the CDP target info settles.
	t.Run("Title", func(t *testing.T) {
		title, err := browser.Title()
		if err != nil {
			t.Fatalf("Title failed: %v", err)
		}
		if title != "Test Page - Forms" {
			t.Errorf("Expected title %q, got %q", "Test Page - Forms", title)
		}
	})

	// 4. Test Observe
	t.Run("Observe", func(t *testing.T) {
		obs, err := browser.Observe()
		if err != nil {
			t.Fatalf("Observe failed: %v", err)
		}

		if len(obs.Forms) == 0 {
			t.Error("Expected at least 1 form")
		}

		// Verify email input exists (either standalone or in form)
		foundEmail := false
		countInputs := len(obs.Inputs)

		for _, input := range obs.Inputs {
			if input.Name == "email" {
				foundEmail = true
			}
		}

		for _, form := range obs.Forms {
			countInputs += len(form.Fields)
			for _, field := range form.Fields {
				if field.Name == "email" {
					foundEmail = true
				}
			}
		}

		if countInputs < 2 {
			t.Errorf("Expected at least 2 inputs (email, password), found %d", countInputs)
		}
		if !foundEmail {
			t.Error("Observe did not find email input")
		}
	})

	// 5. Test Act (Type)
	t.Run("Act_Type", func(t *testing.T) {
		res, err := browser.Act("type test@example.com into email")
		if err != nil {
			t.Fatalf("Act failed: %v", err)
		}
		if !res.Success {
			t.Errorf("Act reported failure: %s", res.Message)
		}

		// Verify value
		valResult, err := browser.Execute(`document.getElementById('email').value`)
		if err != nil {
			t.Fatalf("Failed to verify value: %v", err)
		}

		valStr := fmt.Sprintf("%s", valResult)
		if !strings.Contains(valStr, "test@example.com") {
			t.Errorf("Value not set correctly, got: %s", valStr)
		}
	})

	// 6. Test Act (Click)
	t.Run("Act_Click", func(t *testing.T) {
		// Click the button that changes background to red
		res, err := browser.Act("click the click me button")
		if err != nil {
			t.Fatalf("Act click failed: %v", err)
		}
		if !res.Success {
			t.Errorf("Act click reported failure: %s", res.Message)
		}

		// Allow JS to run
		time.Sleep(100 * time.Millisecond)

		// Verify background color
		bgResult, err := browser.Execute(`document.body.style.backgroundColor`)
		if err != nil {
			t.Fatalf("Failed to verify background: %v", err)
		}

		bgStr := fmt.Sprintf("%s", bgResult)
		if !strings.Contains(bgStr, "red") {
			t.Errorf("Background did not change to red, got: %s", bgStr)
		}
	})
}
