package attributes

// Styleable is implemented by values that support attributes
type Styleable interface {
	Reader
	Writer
}
