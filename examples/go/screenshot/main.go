package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/akashrtd/mime/pkg/mime"
)

func main() {
	ctx := context.Background()

	m, err := mime.New(ctx)
	if err != nil {
		log.Fatalf("Failed to create MIME instance: %v", err)
	}
	defer m.Close()

	// Navigate to a website
	fmt.Println("Navigating to example.com...")
	if err := m.Navigate("https://example.com"); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	// Capture screenshot
	fmt.Println("Capturing screenshot...")
	screenshot, err := m.Screenshot()
	if err != nil {
		log.Fatalf("Failed to capture screenshot: %v", err)
	}

	// Decode base64
	data, err := base64.StdEncoding.DecodeString(screenshot)
	if err != nil {
		log.Fatalf("Failed to decode screenshot: %v", err)
	}

	// Save to file
	if err := os.WriteFile("screenshot.png", data, 0644); err != nil {
		log.Fatalf("Failed to save screenshot: %v", err)
	}

	fmt.Println("Screenshot saved to screenshot.png")
}
