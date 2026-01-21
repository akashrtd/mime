package mime

import (
	"testing"
)

func TestParseInstruction(t *testing.T) {
	tests := []struct {
		input          string
		expectedAction string
		expectedTarget string
	}{
		{"click login", "click", "login"},
		{"click the submit button", "click", "submit button"},
		{"Click on SignUp", "click", "signup"},
		{"type hello into email", "type", "email"},
		{"enter 'my password' into password field", "type", "password field"},
		{"scroll to footer", "scroll", "footer"},
		{"hover over the menu", "hover", "menu"},
		{"submit form", "click", "submit form"}, // default
	}

	for _, tt := range tests {
		action, target := parseInstruction(tt.input)
		if action != tt.expectedAction {
			t.Errorf("parseInstruction(%q) action = %q, want %q", tt.input, action, tt.expectedAction)
		}
		if target != tt.expectedTarget {
			t.Errorf("parseInstruction(%q) target = %q, want %q", tt.input, target, tt.expectedTarget)
		}
	}
}

func TestExtractTextFromInstruction(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"type hello into email", "hello"},
		{"type 'hello world' into description", "hello world"},
		{"enter 123 in zip code", "123"},
		{"click login", ""},
	}

	for _, tt := range tests {
		text := extractTextFromInstruction(tt.input)
		if text != tt.expected {
			t.Errorf("extractTextFromInstruction(%q) = %q, want %q", tt.input, text, tt.expected)
		}
	}
}

func TestFuzzyMatch(t *testing.T) {
	tests := []struct {
		text     string
		target   string
		expected bool
	}{
		{"Login Button", "login", true},
		{"Submit Request", "submit", true},
		{"Submit Request", "request", true},
		{"Sign Up", "signup", false}, // space difference not handled yet
		{"Hello World", "world", true},
		{"Hello World", "foo", false},
	}

	for _, tt := range tests {
		match := fuzzyMatch(tt.text, tt.target)
		if match != tt.expected {
			t.Errorf("fuzzyMatch(%q, %q) = %v, want %v", tt.text, tt.target, match, tt.expected)
		}
	}
}

func TestFindBestMatch(t *testing.T) {
	obs := &PageObservation{
		Clickable: []Clickable{
			{Text: "Login", Selector: "#login-btn"},
			{Text: "Sign Up", Selector: "#signup-btn"},
		},
		Inputs: []InputInfo{
			{Name: "email", Selector: "#email-input"},
			{Placeholder: "Search...", Selector: "#search-input"},
		},
		Forms: []FormInfo{
			{
				Fields: []InputInfo{
					{Name: "password", Selector: "#pwd"},
				},
				Submit: &Clickable{Text: "Submit Form", Selector: "#form-submit"},
			},
		},
	}

	tests := []struct {
		target           string
		expectedSelector string
	}{
		{"login", "#login-btn"},
		{"sign up", "#signup-btn"},
		{"email", "#email-input"},
		{"password", "#pwd"},
		{"search", "#search-input"},
		{"submit form", "#form-submit"},
		{"nonexistent", "text=nonexistent"},
	}

	for _, tt := range tests {
		selector := findBestMatch(obs, tt.target)
		if selector != tt.expectedSelector {
			t.Errorf("findBestMatch(%q) = %q, want %q", tt.target, selector, tt.expectedSelector)
		}
	}
}
