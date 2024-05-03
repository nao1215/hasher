package hasher

import "github.com/reusee/mmh3"

// newMmh3Hasher creates a new Hasher instance for MurmurHash3 algorithm.
func newMmh3Hasher() Hasher {
	return &hasher{HashFunc: mmh3.New128}
}
