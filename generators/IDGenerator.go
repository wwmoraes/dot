package generators

// IDGenerator is implemented by pseudorandom generator values that output
// strings safe enough to be used within a single dot file
type IDGenerator interface {
	String() string
	Stringn(length int) string
}
