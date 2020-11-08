package dot

import (
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// UIDGenerator creates random strings seeded with its instance creation time
type UIDGenerator struct {
	random *rand.Rand
}

// NewUIDGenerator returns a new instance of UIDGenerator seeded with the
// current time
func NewUIDGenerator() *UIDGenerator {
	return &UIDGenerator{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// String generates a random string with given length
func (generator *UIDGenerator) String(length int) string {
	b := make([]byte, length)

	for i := 0; i < length; i++ {
		b[i] = charset[generator.random.Intn(len(charset))]
	}

	return string(b)
}
