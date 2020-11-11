package attributes

// Identity is implemented by values that have an immutable identifier
type Identity interface {
	// ID returns the immutable id
	ID() string
}
