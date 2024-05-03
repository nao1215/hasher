package hasher

import "io"

// Hasher is an interface that contains the methods to generate and compare hashes.
type Hasher interface {
	// GenHashFromString generates a hash from a string.
	GenHashFromString(string) ([]byte, error)
	// GenHashFromIOReader generates a hash from an io.Reader.
	GenHashFromIOReader(io.Reader) ([]byte, error)
	// CmpHashAndString compares a hash and a string.
	// If the hash and the string are the same, nil is returned.
	CmpHashAndString([]byte, string) error
	// CmpHashAndIOReader compares a hash and an io.Reader.
	// If the hash and the io.Reader are the same, nil is returned.
	CmpHashAndIOReader([]byte, io.Reader) error
}
