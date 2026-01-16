package main

import (
	"context"
	"fmt"
	"log"

	"github.com/akashrtd/mime/pkg/mime"
)

func main() {
	ctx := context.Background()

	// Create MIME instance
	m, err := mime.New(ctx)
	if err != nil {
		log.Fatalf("Failed to create MIME instance: %v", err)
	}
	defer m.Close()

	// Navigate to Hacker News
	fmt.Println("Navigating to Hacker News...")
	if err := m.Navigate("https://news.ycombinator.com"); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	// Wait for page to load
	if err := m.WaitFor(".titleline"); err != nil {
		log.Fatalf("Failed to wait for element: %v", err)
	}

	// Extract top story title
	title, err := m.Extract(".titleline a")
	if err != nil {
		log.Fatalf("Failed to extract title: %v", err)
	}

	fmt.Printf("Top story: %s\n", title)

	// Get page URL
	url := m.URL()
	fmt.Printf("Current URL: %s\n", url)

	// Get page title
	pageTitle, err := m.Title()
	if err != nil {
		log.Fatalf("Failed to get page title: %v", err)
	}
	fmt.Printf("Page title: %s\n", pageTitle)
}
