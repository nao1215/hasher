package hasher

import (
	"hash/adler32"
)

func newAdler32Hasher() Hasher {
	return &hasher32{HashFunc: adler32.New}
}
