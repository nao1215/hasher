package hasher

import (
	"bytes"
	"crypto/md5" //nolint:gosec
	"io"
)

type md5sumHasher struct{}

// GenHashFromString generates a hash from a string using the md5sum algorithm.
func (m *md5sumHasher) GenHashFromString(s string) ([]byte, error) {
	h := md5.New() //nolint:gosec
	if _, err := h.Write([]byte(s)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// GenHashFromIOReader generates a hash from an io.Reader using the md5sum algorithm.
func (m *md5sumHasher) GenHashFromIOReader(r io.Reader) ([]byte, error) {
	h := md5.New() //nolint:gosec
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// CmpHashAndString compares a hash and a string using the md5sum algorithm.
func (m *md5sumHasher) CmpHashAndString(hashA []byte, s string) error {
	hashB, err := m.GenHashFromString(s)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}

// CmpHashAndIOReader compares a hash and an io.Reader using the md5sum algorithm.
func (m *md5sumHasher) CmpHashAndIOReader(hashA []byte, r io.Reader) error {
	hashB, err := m.GenHashFromIOReader(r)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}
