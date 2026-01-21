package mime

import (
	"strings"
	"testing"
)

func TestHtmlToMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string // contains (because library might add newlines)
	}{
		{
			name:     "Header",
			html:     "<h1>Hello World</h1>",
			expected: "# Hello World",
		},
		{
			name:     "Link",
			html:     `<a href="https://example.com">Example</a>`,
			expected: "[Example](https://example.com)",
		},
		{
			name:     "Formatting",
			html:     "<p><b>Bold</b> and <i>Italic</i></p>",
			expected: "**Bold** and *Italic*",
		},
		{
			name:     "List",
			html:     "<ul><li>Item 1</li><li>Item 2</li></ul>",
			expected: "- Item 1", // Expect partial match
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md, err := htmlToMarkdown(tt.html)
			if err != nil {
				t.Fatalf("htmlToMarkdown failed: %v", err)
			}

			if !strings.Contains(md, tt.expected) {
				t.Errorf("markdown %q should contain %q", md, tt.expected)
			}
		})
	}
}
