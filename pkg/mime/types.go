package mime

// WaitForArgs arguments for wait_for tool
type WaitForArgs struct {
	Selector string `json:"selector"`
}

// ScrollArgs arguments for scroll tool
type ScrollArgs struct {
	Selector string `json:"selector,omitempty"`
	X        int    `json:"x,omitempty"`
	Y        int    `json:"y,omitempty"`
}

// HoverArgs arguments for hover tool
type HoverArgs struct {
	Selector string `json:"selector"`
}

// MarkdownInput arguments for markdown tool
type MarkdownInput struct {
	FullPage bool `json:"full_page,omitempty"`
}

// MarkdownOutput output for markdown tool
type MarkdownOutput struct {
	Markdown string `json:"markdown"`
	URL      string `json:"url"`
}

// MetadataOutput output for metadata tool
type MetadataOutput struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	URL         string            `json:"url"`
	Canonical   string            `json:"canonical,omitempty"`
	OG          map[string]string `json:"og,omitempty"`
}

// LinksOutput output for links tool
type LinksOutput struct {
	Links []Link `json:"links"`
	Count int    `json:"count"`
}

// CookieInfo simplified cookie information
type CookieInfo struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Domain string `json:"domain"`
	Path   string `json:"path"`
}

// CookiesOutput output for get_cookies tool
type CookiesOutput struct {
	Cookies []CookieInfo `json:"cookies"`
	Count   int          `json:"count"`
}

// StatusOutput generic status output
type StatusOutput struct {
	Status string `json:"status"`
}

// ObserveOutput output for observe tool
type ObserveOutput struct {
	URL       string      `json:"url"`
	Title     string      `json:"title"`
	Forms     []FormInfo  `json:"forms"`
	Clickable []Clickable `json:"clickable"`
	Inputs    []InputInfo `json:"inputs"`
	Content   ContentInfo `json:"content"`
}

// ActInput input for act tool
type ActInput struct {
	Instruction string `json:"instruction"`
}

// ActOutput output for act tool
type ActOutput struct {
	Success bool   `json:"success"`
	Action  string `json:"action"`
	Target  string `json:"target"`
	Message string `json:"message,omitempty"`
}
