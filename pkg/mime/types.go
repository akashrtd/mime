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

// types.go contains shared type definitions for the MIME package
// This file can be extended with common types and constants as needed
