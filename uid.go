package dot

import (
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// UIDGenerator creates random strings seeded with its instance creation time
type UIDGenerator struct {
	random *rand.Rand
	length int
}

// NewUIDGenerator returns a new instance of UIDGenerator seeded with the
// current time
func NewUIDGenerator(length int) *UIDGenerator {
	return &UIDGenerator{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
		length: length,
	}
}

// String generates a random string with the default length
func (generator *UIDGenerator) String() string {
	b := make([]byte, generator.length)

	for i := 0; i < generator.length; i++ {
		b[i] = charset[generator.random.Intn(len(charset))]
	}

	return string(b)
}

// Stringn generates a random string with given length
func (generator *UIDGenerator) Stringn(length int) string {
	b := make([]byte, length)

	for i := 0; i < length; i++ {
		b[i] = charset[generator.random.Intn(len(charset))]
	}

	return string(b)
}
