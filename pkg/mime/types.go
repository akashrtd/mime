package mime

// WaitForArgs arguments for wait_for tool
type WaitForArgs struct {
	Selector string `json:"selector" jsonschema:"description=CSS selector or text=Query to wait for"`
}

// ScrollArgs arguments for scroll tool
type ScrollArgs struct {
	Selector string `json:"selector,omitempty" jsonschema:"description=Optional. Selector to scroll into view"`
	X        int    `json:"x,omitempty" jsonschema:"description=Pixels to scroll horizontally (if selector not provided)"`
	Y        int    `json:"y,omitempty" jsonschema:"description=Pixels to scroll vertically (if selector not provided)"`
}

// HoverArgs arguments for hover tool
type HoverArgs struct {
	Selector string `json:"selector" jsonschema:"description=CSS selector or text=Query to hover over"`
}

// types.go contains shared type definitions for the MIME package
// This file can be extended with common types and constants as needed
