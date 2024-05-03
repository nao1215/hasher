package hasher

import "github.com/jzelinskie/whirlpool"

// newWhirlpoolHasher creates a new Hasher instance for Whirlpool algorithm.
func newWhirlpoolHasher() Hasher {
	return &hasher{HashFunc: whirlpool.New}
}
