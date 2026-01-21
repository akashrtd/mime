package mime

import (
	"fmt"
	"net/url"
)

// MapResult result of a map operation
type MapResult struct {
	URLs  []string `json:"urls"`
	Total int      `json:"total"`
}

// Map discovers all unique links on a site up to a certain limit
// unlike Crawl, it focuses on URL discovery and site structure
func (b *Browser) Map(startURL string) (*MapResult, error) {
	// For now, map is essentially a simplified crawl that only cares about URLs
	// We'll limit it to same domain and max 50 pages for speed

	startURLObj, err := url.Parse(startURL)
	if err != nil {
		return nil, fmt.Errorf("invalid start URL: %w", err)
	}

	visited := make(map[string]bool)
	queue := []string{startURL}
	var foundURLs []string

	// Max limit to avoid infinite loops
	const maxPages = 50

	visited[startURL] = true
	foundURLs = append(foundURLs, startURL)

	for len(queue) > 0 && len(visited) < maxPages {
		currentURL := queue[0]
		queue = queue[1:]

		// Navigate
		if err := b.Navigate(currentURL); err != nil {
			continue
		}

		// Get all links
		links, err := b.Links()
		if err != nil {
			continue
		}

		for _, link := range links {
			linkURL := link.URL

			// Skip if already visited/queued
			if visited[linkURL] {
				continue
			}

			// Same domain check
			u, err := url.Parse(linkURL)
			if err == nil && u.Host == startURLObj.Host {
				visited[linkURL] = true
				foundURLs = append(foundURLs, linkURL)
				queue = append(queue, linkURL)
			}
		}
	}

	return &MapResult{
		URLs:  foundURLs,
		Total: len(foundURLs),
	}, nil
}
