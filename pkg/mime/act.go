package mime

import (
	"fmt"
	"strings"
)

// ActResult represents the result of a natural language action
type ActResult struct {
	Success bool   `json:"success"`
	Action  string `json:"action"`
	Target  string `json:"target"`
	Message string `json:"message,omitempty"`
}

// Act performs an action based on natural language instruction
// Examples: "click the login button", "type hello into the email field"
func (b *Browser) Act(instruction string) (*ActResult, error) {
	// 1. Observe the page
	obs, err := b.Observe()
	if err != nil {
		return nil, fmt.Errorf("failed to observe page: %w", err)
	}

	// 2. Parse the instruction
	action, target := parseInstruction(instruction)

	// 3. Find matching element
	selector := findBestMatch(obs, target)
	if selector == "" {
		return &ActResult{
			Success: false,
			Action:  action,
			Message: fmt.Sprintf("Could not find element matching: %s", target),
		}, nil
	}

	// 4. Perform the action
	var actionErr error
	switch action {
	case "click":
		actionErr = b.Click(selector)
	case "type":
		text := extractTextFromInstruction(instruction)
		if text != "" {
			actionErr = b.Type(selector, text)
		} else {
			return &ActResult{
				Success: false,
				Action:  action,
				Target:  selector,
				Message: "No text specified for type action",
			}, nil
		}
	case "scroll":
		actionErr = b.Scroll(selector, 0, 0)
	case "hover":
		actionErr = b.Hover(selector)
	default:
		// Default to click
		actionErr = b.Click(selector)
	}

	if actionErr != nil {
		return &ActResult{
			Success: false,
			Action:  action,
			Target:  selector,
			Message: actionErr.Error(),
		}, nil
	}

	return &ActResult{
		Success: true,
		Action:  action,
		Target:  selector,
	}, nil
}

// parseInstruction extracts action and target from natural language
func parseInstruction(instruction string) (action, target string) {
	instruction = strings.ToLower(strings.TrimSpace(instruction))

	// Handle "click on X" or "click the X" or "click X"
	if strings.HasPrefix(instruction, "click ") {
		target = strings.TrimPrefix(instruction, "click ")
		target = strings.TrimPrefix(target, "on ")
		target = strings.TrimPrefix(target, "the ")
		return "click", target
	}

	// Handle "type X into Y" or "enter X in Y"
	if strings.HasPrefix(instruction, "type ") || strings.HasPrefix(instruction, "enter ") {
		instruction = strings.TrimPrefix(instruction, "type ")
		instruction = strings.TrimPrefix(instruction, "enter ")

		// Look for "into" or "in" to find the target
		for _, sep := range []string{" into ", " in ", " into the ", " in the "} {
			if idx := strings.Index(instruction, sep); idx != -1 {
				target = strings.TrimSpace(instruction[idx+len(sep):])
				return "type", target
			}
		}
		return "type", instruction
	}

	// Handle "scroll to X"
	if strings.HasPrefix(instruction, "scroll ") {
		target = strings.TrimPrefix(instruction, "scroll ")
		target = strings.TrimPrefix(target, "to ")
		target = strings.TrimPrefix(target, "the ")
		return "scroll", target
	}

	// Handle "hover over X"
	if strings.HasPrefix(instruction, "hover ") {
		target = strings.TrimPrefix(instruction, "hover ")
		target = strings.TrimPrefix(target, "over ")
		target = strings.TrimPrefix(target, "the ")
		return "hover", target
	}

	// Default: treat as click target
	return "click", instruction
}

// extractTextFromInstruction extracts the text to type from an instruction
func extractTextFromInstruction(instruction string) string {
	instruction = strings.ToLower(instruction)

	// Pattern: "type X into Y" or "enter X in Y"
	for _, prefix := range []string{"type ", "enter "} {
		if strings.HasPrefix(instruction, prefix) {
			rest := strings.TrimPrefix(instruction, prefix)

			// Find where the target description starts
			for _, sep := range []string{" into ", " in "} {
				if idx := strings.Index(rest, sep); idx != -1 {
					text := rest[:idx]
					// Remove quotes if present
					text = strings.Trim(text, "\"'")
					return text
				}
			}
		}
	}

	return ""
}

// findBestMatch finds the best matching selector from the observation
func findBestMatch(obs *PageObservation, target string) string {
	target = strings.ToLower(strings.TrimSpace(target))

	// Remove common words
	target = strings.ReplaceAll(target, "button", "")
	target = strings.ReplaceAll(target, "link", "")
	target = strings.ReplaceAll(target, "field", "")
	target = strings.ReplaceAll(target, "input", "")
	target = strings.TrimSpace(target)

	// Check clickable elements (buttons, links)
	for _, c := range obs.Clickable {
		if fuzzyMatch(c.Text, target) {
			return c.Selector
		}
	}

	// Check form fields
	for _, f := range obs.Forms {
		for _, field := range f.Fields {
			if fuzzyMatch(field.Name, target) || fuzzyMatch(field.Placeholder, target) {
				return field.Selector
			}
		}
		// Check form submit button
		if f.Submit != nil && fuzzyMatch(f.Submit.Text, target) {
			return f.Submit.Selector
		}
	}

	// Check standalone inputs
	for _, i := range obs.Inputs {
		if fuzzyMatch(i.Name, target) || fuzzyMatch(i.Placeholder, target) {
			return i.Selector
		}
	}

	// Try text= selector as fallback
	if len(target) > 2 {
		return "text=" + target
	}

	return ""
}

// fuzzyMatch checks if text contains the target (case insensitive)
func fuzzyMatch(text, target string) bool {
	if target == "" {
		return false
	}
	text = strings.ToLower(text)
	target = strings.ToLower(target)

	// Exact match
	if text == target {
		return true
	}

	// Contains match
	if strings.Contains(text, target) {
		return true
	}

	// Check if all words in target appear in text
	targetWords := strings.Fields(target)
	for _, word := range targetWords {
		if len(word) > 2 && strings.Contains(text, word) {
			return true
		}
	}

	return false
}
