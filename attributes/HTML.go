package attributes

// HTML renders the content as HTML, as supported by some attributes (e.g. label)
type HTML struct {
	value string
}

// NewHTML creates a new HTML string, that'll keep HTML code intact
func NewHTML(value string) *HTML {
	return &HTML{value}
}

// String returns the HTML string
func (data HTML) String() string {
	return data.value
}
