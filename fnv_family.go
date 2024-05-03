package hasher

import (
	"bytes"
	"hash/fnv"
	"io"
)

// newFnv128Hasher creates a new Hasher instance for FNV-128 algorithm.
func newFnv128Hasher() Hasher {
	return &hasher{HashFunc: fnv.New128}
}

// newFnv128aHasher creates a new Hasher instance for FNV-128a algorithm.
func newFnv128aHasher() Hasher {
	return &hasher{HashFunc: fnv.New128a}
}

type fnv32Hasher struct{}

// GenHashFromString generates a hash from a string using the FNV-32 algorithm.
func (f *fnv32Hasher) GenHashFromString(s string) ([]byte, error) {
	h := fnv.New32()
	if _, err := h.Write([]byte(s)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// GenHashFromIOReader generates a hash from an io.Reader using the FNV-32 algorithm.
func (f *fnv32Hasher) GenHashFromIOReader(r io.Reader) ([]byte, error) {
	h := fnv.New32()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// CmpHashAndString compares a hash and a string using the FNV-32 algorithm.
func (f *fnv32Hasher) CmpHashAndString(hashA []byte, s string) error {
	hashB, err := f.GenHashFromString(s)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}

// CmpHashAndIOReader compares a hash and an io.Reader using the FNV-32 algorithm.
func (f *fnv32Hasher) CmpHashAndIOReader(hashA []byte, r io.Reader) error {
	hashB, err := f.GenHashFromIOReader(r)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}

type fnv32aHasher struct{}

// GenHashFromString generates a hash from a string using the FNV-32a algorithm.
func (f *fnv32aHasher) GenHashFromString(s string) ([]byte, error) {
	h := fnv.New32a()
	if _, err := h.Write([]byte(s)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// GenHashFromIOReader generates a hash from an io.Reader using the FNV-32a algorithm.
func (f *fnv32aHasher) GenHashFromIOReader(r io.Reader) ([]byte, error) {
	h := fnv.New32a()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// CmpHashAndString compares a hash and a string using the FNV-32 algorithm.
func (f *fnv32aHasher) CmpHashAndString(hashA []byte, s string) error {
	hashB, err := f.GenHashFromString(s)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}

// CmpHashAndIOReader compares a hash and an io.Reader using the FNV-32 algorithm.
func (f *fnv32aHasher) CmpHashAndIOReader(hashA []byte, r io.Reader) error {
	hashB, err := f.GenHashFromIOReader(r)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}

type fnv64Hasher struct{}

// GenHashFromString generates a hash from a string using the FNV-64 algorithm.
func (f *fnv64Hasher) GenHashFromString(s string) ([]byte, error) {
	h := fnv.New64()
	if _, err := h.Write([]byte(s)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// GenHashFromIOReader generates a hash from an io.Reader using the FNV-64 algorithm.
func (f *fnv64Hasher) GenHashFromIOReader(r io.Reader) ([]byte, error) {
	h := fnv.New64()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// CmpHashAndString compares a hash and a string using the FNV-64 algorithm.
func (f *fnv64Hasher) CmpHashAndString(hashA []byte, s string) error {
	hashB, err := f.GenHashFromString(s)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}

// CmpHashAndIOReader compares a hash and an io.Reader using the FNV-64 algorithm.
func (f *fnv64Hasher) CmpHashAndIOReader(hashA []byte, r io.Reader) error {
	hashB, err := f.GenHashFromIOReader(r)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}

type fnv64aHasher struct{}

// GenHashFromString generates a hash from a string using the FNV-64a algorithm.
func (f *fnv64aHasher) GenHashFromString(s string) ([]byte, error) {
	h := fnv.New64a()
	if _, err := h.Write([]byte(s)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// GenHashFromIOReader generates a hash from an io.Reader using the FNV-64a algorithm.
func (f *fnv64aHasher) GenHashFromIOReader(r io.Reader) ([]byte, error) {
	h := fnv.New64a()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// CmpHashAndString compares a hash and a string using the FNV-64a algorithm.
func (f *fnv64aHasher) CmpHashAndString(hashA []byte, s string) error {
	hashB, err := f.GenHashFromString(s)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}

// CmpHashAndIOReader compares a hash and an io.Reader using the FNV-64a algorithm.
func (f *fnv64aHasher) CmpHashAndIOReader(hashA []byte, r io.Reader) error {
	hashB, err := f.GenHashFromIOReader(r)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}
