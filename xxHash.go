package hasher

import (
	"github.com/cespare/xxhash"
)

// newXXHasher creates a new Hasher instance for XXHash algorithm.
func newXXHasher() Hasher {
	return &hasher64{HashFunc: xxhash.New}
}
