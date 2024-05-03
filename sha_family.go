package hasher

import (
	"bytes"
	"crypto/sha1" //nolint:gosec
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"io"
)

// shaHasher represents a generic hasher for SHA algorithms.
type shaHasher struct {
	HashFunc func() hash.Hash
}

// GenHashFromString generates a hash from a string using the specified hash function.
// Supported hash functions: SHA-1, SHA-256, SHA-512.
func (s *shaHasher) GenHashFromString(str string) ([]byte, error) {
	h := s.HashFunc()
	if _, err := h.Write([]byte(str)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// GenHashFromIOReader generates a hash from an io.Reader using the specified hash function.
// Supported hash functions: SHA-1, SHA-256, SHA-512.
func (s *shaHasher) GenHashFromIOReader(r io.Reader) ([]byte, error) {
	h := s.HashFunc()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// CmpHashAndString compares a hash and a string using the specified hash function.
func (s *shaHasher) CmpHashAndString(hashA []byte, str string) error {
	hashB, err := s.GenHashFromString(str)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}

// CmpHashAndIOReader compares a hash and an io.Reader using the specified hash function.
// Supported hash functions: SHA-1, SHA-256, SHA-512.
func (s *shaHasher) CmpHashAndIOReader(hashA []byte, r io.Reader) error {
	hashB, err := s.GenHashFromIOReader(r)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}

// newSHA1Hasher creates a new Hasher instance for SHA-1 algorithm.
func newSHA1Hasher() Hasher {
	return &shaHasher{HashFunc: sha1.New}
}

// newSHA256Hasher creates a new Hasher instance for SHA-256 algorithm.
func newSHA256Hasher() Hasher {
	return &shaHasher{HashFunc: sha256.New}
}

// newSHA512Hasher creates a new Hasher instance for SHA-512 algorithm.
func newSHA512Hasher() Hasher {
	return &shaHasher{HashFunc: sha512.New}
}
