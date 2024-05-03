package hasher

import "hash/crc32"

// newCRC32Hasher creates a new Hasher instance for CRC32 algorithm.
func newCRC32Hasher() Hasher {
	return &hasher32{HashFunc: crc32.NewIEEE}
}
