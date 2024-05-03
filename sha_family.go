package hasher

import (
	"crypto/sha1" //nolint:gosec
	"crypto/sha256"
	"crypto/sha512"
)

// newSHA1Hasher creates a new Hasher instance for SHA-1 algorithm.
func newSHA1Hasher() Hasher {
	return &hasher{HashFunc: sha1.New}
}

// newSHA256Hasher creates a new Hasher instance for SHA-256 algorithm.
func newSHA256Hasher() Hasher {
	return &hasher{HashFunc: sha256.New}
}

// newSHA512Hasher creates a new Hasher instance for SHA-512 algorithm.
func newSHA512Hasher() Hasher {
	return &hasher{HashFunc: sha512.New}
}
