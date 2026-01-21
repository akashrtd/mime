package mime

import (
	"encoding/json"
	"fmt"
)

// PageObservation provides structured page understanding for AI agents
type PageObservation struct {
	URL       string      `json:"url"`
	Title     string      `json:"title"`
	Forms     []FormInfo  `json:"forms"`
	Clickable []Clickable `json:"clickable"`
	Inputs    []InputInfo `json:"inputs"`
	Content   ContentInfo `json:"content"`
}

// FormInfo describes a form on the page
type FormInfo struct {
	ID     string      `json:"id,omitempty"`
	Action string      `json:"action"`
	Method string      `json:"method"`
	Fields []InputInfo `json:"fields"`
	Submit *Clickable  `json:"submit,omitempty"`
}

// Clickable describes an interactive element
type Clickable struct {
	Text     string `json:"text"`
	Selector string `json:"selector"`
	Type     string `json:"type"` // "button", "link", "submit"
}

// InputInfo describes an input field
type InputInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Selector    string `json:"selector"`
	Placeholder string `json:"placeholder,omitempty"`
	Required    bool   `json:"required"`
}

// ContentInfo summarizes page content
type ContentInfo struct {
	Headings   []string `json:"headings"`
	Paragraphs []string `json:"paragraphs"`
	ImageCount int      `json:"image_count"`
}

// Observe analyzes the page and returns structured information for AI agents
func (b *Browser) Observe() (*PageObservation, error) {
	script := `() => {
		// Helper to generate a unique selector
		const getSelector = (el) => {
			if (el.id) return '#' + el.id;
			if (el.name) return '[name="' + el.name + '"]';
			if (el.className && typeof el.className === 'string') {
				const classes = el.className.trim().split(/\s+/).slice(0, 2).join('.');
				if (classes) return el.tagName.toLowerCase() + '.' + classes;
			}
			return el.tagName.toLowerCase();
		};

		// Extract forms
		const forms = [...document.forms].slice(0, 10).map(f => {
			const fields = [...f.elements]
				.filter(e => e.tagName !== 'BUTTON' && e.type !== 'submit' && e.type !== 'hidden')
				.slice(0, 20)
				.map(e => ({
					name: e.name || '',
					type: e.type || 'text',
					selector: getSelector(e),
					placeholder: e.placeholder || '',
					required: e.required || false
				}));
			
			const submitBtn = f.querySelector('button[type="submit"], input[type="submit"], button:not([type])');
			const submit = submitBtn ? {
				text: (submitBtn.textContent || submitBtn.value || 'Submit').trim().slice(0, 50),
				selector: getSelector(submitBtn),
				type: 'submit'
			} : null;

			return {
				id: f.id || '',
				action: f.action || '',
				method: (f.method || 'get').toUpperCase(),
				fields,
				submit
			};
		});

		// Extract clickable elements (buttons, links)
		const clickable = [...document.querySelectorAll('a[href], button, [role="button"], input[type="submit"], input[type="button"]')]
			.filter(e => {
				const rect = e.getBoundingClientRect();
				return rect.width > 0 && rect.height > 0 && e.offsetParent !== null;
			})
			.slice(0, 30)
			.map(e => ({
				text: (e.textContent || e.value || e.title || e.getAttribute('aria-label') || '').trim().slice(0, 100),
				selector: getSelector(e),
				type: e.tagName.toLowerCase() === 'a' ? 'link' : 'button'
			}))
			.filter(c => c.text.length > 0);

		// Extract standalone inputs (not in forms)
		const formInputs = new Set([...document.forms].flatMap(f => [...f.elements]));
		const inputs = [...document.querySelectorAll('input, textarea, select')]
			.filter(e => !formInputs.has(e) && e.type !== 'hidden')
			.slice(0, 20)
			.map(e => ({
				name: e.name || '',
				type: e.type || 'text',
				selector: getSelector(e),
				placeholder: e.placeholder || '',
				required: e.required || false
			}));

		// Extract content summary
		const headings = [...document.querySelectorAll('h1, h2, h3')]
			.slice(0, 10)
			.map(h => h.textContent.trim().slice(0, 200))
			.filter(h => h.length > 0);

		const paragraphs = [...document.querySelectorAll('p')]
			.slice(0, 5)
			.map(p => p.textContent.trim().slice(0, 300))
			.filter(p => p.length > 20);

		return {
			url: location.href,
			title: document.title,
			forms,
			clickable,
			inputs,
			content: {
				headings,
				paragraphs,
				image_count: document.images.length
			}
		};
	}`

	result, err := b.page.Eval(script)
	if err != nil {
		return nil, fmt.Errorf("failed to observe page: %w", err)
	}

	jsonBytes, err := result.Value.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal observation: %w", err)
	}

	var obs PageObservation
	if err := json.Unmarshal(jsonBytes, &obs); err != nil {
		return nil, fmt.Errorf("failed to parse observation: %w", err)
	}

	return &obs, nil
}
