package attributes

// Object is implemented by values that support attributes and must be unique
type Object interface {
	Reader
	Writer
	// ID returns the immutable id
	ID() string
}
