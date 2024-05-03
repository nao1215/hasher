// Package hasher provides functionality for generating and comparing hashes using various algorithms.
//
// The Hash struct contains methods for generating and comparing hashes. By default, the MD5 algorithm is used,
// but the user can specify a different algorithm using options.
//
// Example usage:
//
//	// Create a new Hash instance with default options.
//	h := hasher.NewHash()
//
//	// Generate a hash from a string.
//	hash, err := h.Generate("example")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Compare a hash with a string.
//	err := h.Compare(hash, "example")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Generate a hash from an io.Reader.
//	file, err := os.Open("example.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer file.Close()
//	hash, err = h.Generate(file)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Compare a hash with an io.Reader.
//	file, err = os.Open("example.txt")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer file.Close()
//	err = h.Compare(hash, file)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Supported algorithms:
//   - MD5
//   - SHA-1
//   - SHA-256
//   - SHA-512
package hasher

import (
	"fmt"
	"io"
)

// Hash is a struct that contains the methods to generate and compare hashes.
type Hash struct {
	hasher Hasher
}

// NewHash returns a new Hasher struct. Default hash algorithm is MD5SUM.
// The user can specify a different algorithm using options.
// e.g. NewHash(WithSha1Algorithm())
func NewHash(opts ...Option) *Hash {
	h := &Hash{
		hasher: &md5sumHasher{},
	}

	for _, opt := range opts {
		opt(h)
	}
	return h
}

// Generate generates a hash from the input.
// The input can be a string or an io.Reader. If the input is not a string or an io.Reader,
// ErrUnsupportedInputType is returned.
func (h *Hash) Generate(input any) ([]byte, error) {
	switch v := input.(type) {
	case string:
		return h.hasher.GenHashFromString(v)
	case io.Reader:
		return h.hasher.GenHashFromIOReader(v)
	default:
		return nil, fmt.Errorf("%w: %T", ErrUnsupportedInputType, v)
	}
}

// Compare compares hash and input.
// The input can be a string or an io.Reader. If the input is not a string or an io.Reader, ErrUnsupportedInputType is returned.
// If the hash and the input are the same, nil is returned.
// If the hash and the input are different with hasher support algorithm, an ErrHashMismatch is returned.
func (h *Hash) Compare(hash []byte, input any) error {
	switch v := input.(type) {
	case string:
		return h.hasher.CmpHashAndString(hash, v)
	case io.Reader:
		return h.hasher.CmpHashAndIOReader(hash, v)
	default:
		return fmt.Errorf("%w: %T", ErrUnsupportedInputType, v)
	}
}
