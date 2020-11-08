package attributes

// String plain string that'll be safely escaped
type String struct {
	value string
}

// NewString creates a standard string that'll be safely escaped
func NewString(value string) *String {
	return &String{value}
}

// String returns the string value
func (data String) String() string {
	return data.value
}
