package mime

import (
	"fmt"
	"net/url"
	"strings"
)

// CrawlOptions configuration for crawling
type CrawlOptions struct {
	URL        string   `json:"url"` // Starting URL
	MaxPages   int      `json:"max_pages"`
	MaxDepth   int      `json:"max_depth"`
	SameDomain bool     `json:"same_domain"`
	Patterns   []string `json:"patterns,omitempty"` // simple string match
	Excludes   []string `json:"excludes,omitempty"`
}

// CrawlResult results from a crawl
type CrawlResult struct {
	Pages []PageResult `json:"pages"`
	Total int          `json:"total"`
}

// PageResult result for a single page
type PageResult struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	Markdown string `json:"markdown,omitempty"`
	Status   string `json:"status"` // "success", "error", "skipped"
}

// Crawl visits pages recursively
func (b *Browser) Crawl(startURL string, opts *CrawlOptions) (*CrawlResult, error) {
	if opts == nil {
		opts = &CrawlOptions{
			MaxPages:   10,
			MaxDepth:   2,
			SameDomain: true,
		}
	}
	if opts.MaxPages == 0 {
		opts.MaxPages = 5
	}

	startURLObj, err := url.Parse(startURL)
	if err != nil {
		return nil, fmt.Errorf("invalid start URL: %w", err)
	}

	visited := make(map[string]bool)
	results := []PageResult{}

	// Track depth roughly by queue layers, but for simplicity here we just use max pages constraint mostly
	// To track depth properly we'd need a queue item struct { url, depth }
	type queueItem struct {
		url   string
		depth int
	}

	q := []queueItem{{url: startURL, depth: 0}}
	visited[startURL] = true

	for len(q) > 0 && len(results) < opts.MaxPages {
		current := q[0]
		q = q[1:]

		if current.depth > opts.MaxDepth {
			continue
		}

		// Navigate
		err := b.Navigate(current.url)
		result := PageResult{
			URL: current.url,
		}

		if err != nil {
			result.Status = "error"
			// Log error but continue
			results = append(results, result)
			continue
		}

		// Wait network idle? Or just wait a bit is handled by Navigate usually but mostly it waits for load event

		// Extract content
		title, _ := b.Title()
		result.Title = title

		md, err := b.Markdown(nil)
		if err == nil {
			result.Markdown = md
			result.Status = "success"
		} else {
			result.Status = "partial_error"
		}

		results = append(results, result)

		// If we haven't reached depth limit, get links
		if current.depth < opts.MaxDepth && len(results) < opts.MaxPages {
			links, err := b.Links()
			if err == nil {
				for _, link := range links {
					linkURL := link.URL

					// Filter
					if visited[linkURL] {
						continue
					}

					// Domain check
					if opts.SameDomain {
						u, err := url.Parse(linkURL)
						if err == nil && u.Host != startURLObj.Host {
							continue
						}
					}

					// Pattern check
					if !shouldVisit(linkURL, opts) {
						continue
					}

					visited[linkURL] = true
					q = append(q, queueItem{url: linkURL, depth: current.depth + 1})
				}
			}
		}
	}

	return &CrawlResult{
		Pages: results,
		Total: len(results),
	}, nil
}

func shouldVisit(urlStr string, opts *CrawlOptions) bool {
	// Check excludes
	for _, excl := range opts.Excludes {
		if strings.Contains(urlStr, excl) {
			return false
		}
	}

	// Check includes (if empty, include all)
	if len(opts.Patterns) == 0 {
		return true
	}

	for _, pat := range opts.Patterns {
		if strings.Contains(urlStr, pat) {
			return true
		}
	}

	return false
}
