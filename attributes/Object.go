package attributes

// Object represents a dot object that supports attributes
type Object interface {
	Reader
	Writer
	ID() string
}
