package hasher

import (
	"bytes"
	"io"

	"lukechampine.com/blake3"
)

type blake3Hasher struct{}

// GenHashFromString generates a hash from a string using the blake3 algorithm.
// The hash length is 64 bytes.
func (b *blake3Hasher) GenHashFromString(s string) ([]byte, error) {
	h := blake3.New(64, nil)
	if _, err := h.Write([]byte(s)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// GenHashFromIOReader generates a hash from an io.Reader using the blake3 algorithm.
// The hash length is 64 bytes.
func (b *blake3Hasher) GenHashFromIOReader(r io.Reader) ([]byte, error) {
	h := blake3.New(64, nil)
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// CmpHashAndString compares a hash and a string using the blake3 algorithm.
func (b *blake3Hasher) CmpHashAndString(hashA []byte, s string) error {
	hashB, err := b.GenHashFromString(s)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}

// CmpHashAndIOReader compares a hash and an io.Reader using the blake3 algorithm.
func (b *blake3Hasher) CmpHashAndIOReader(hashA []byte, r io.Reader) error {
	hashB, err := b.GenHashFromIOReader(r)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}
