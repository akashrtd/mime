package mime

import (
	"net/url"
	"strings"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/go-shiori/go-readability"
)

// MarkdownOptions configures markdown extraction
type MarkdownOptions struct {
	OnlyMainContent bool // Extract only main article content
}

// Markdown returns the page content as clean markdown
func (b *Browser) Markdown(opts *MarkdownOptions) (string, error) {
	if opts == nil {
		opts = &MarkdownOptions{OnlyMainContent: true}
	}

	html, err := b.HTML()
	if err != nil {
		return "", err
	}

	pageURL := b.URL()

	if opts.OnlyMainContent {
		// Use readability to extract main content
		parsedURL, err := url.Parse(pageURL)
		if err != nil {
			return htmlToMarkdown(html)
		}

		article, err := readability.FromReader(strings.NewReader(html), parsedURL)
		if err != nil {
			// Fallback to full HTML
			return htmlToMarkdown(html)
		}

		return htmlToMarkdown(article.Content)
	}

	return htmlToMarkdown(html)
}

// htmlToMarkdown converts HTML string to markdown
func htmlToMarkdown(html string) (string, error) {
	md, err := htmltomarkdown.ConvertString(html)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(md), nil
}
