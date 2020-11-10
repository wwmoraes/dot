package attributes

// Identity is implemented by values that can have an identifier
type Identity interface {
	// ID returns the immutable id
	ID() string
}
