package generators

import (
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// randTimeIDGenerator creates random strings seeded with the instance monotonic
// creation time
type randTimeIDGenerator struct {
	random *rand.Rand
	length int
}

// NewRandTimeIDGenerator returns a new instance of randTimeIDGenerator
func NewRandTimeIDGenerator(length int) IDGenerator {
	return &randTimeIDGenerator{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
		length: length,
	}
}

// String generates a random string with the default length
func (generator *randTimeIDGenerator) String() string {
	b := make([]byte, generator.length)

	for i := 0; i < generator.length; i++ {
		b[i] = charset[generator.random.Intn(len(charset))]
	}

	return string(b)
}

// Stringn generates a random string with given length
func (generator *randTimeIDGenerator) Stringn(length int) string {
	b := make([]byte, length)

	for i := 0; i < length; i++ {
		b[i] = charset[generator.random.Intn(len(charset))]
	}

	return string(b)
}
