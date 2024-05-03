package hasher

import (
	"bytes"
	"encoding/binary"
	"image"
	"io"

	"github.com/azr/phash"
)

type pHasher struct{}

// GenHashFromString always returns ErrPhashNotSupportedString because perceptual hashing  does not support string input.
func (p *pHasher) GenHashFromString(_ string) ([]byte, error) {
	return nil, ErrPhashNotSupportedString
}

// CmpHashAndString always returns ErrPhashNotSupportedString because perceptual hashing  does not support string input.
func (p *pHasher) CmpHashAndString(_ []byte, _ string) error {
	return ErrPhashNotSupportedString
}

// GenHashFromIOReader generates a hash from an io.Reader using the perceptual hashing  algorithm.
func (p *pHasher) GenHashFromIOReader(r io.Reader) ([]byte, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	hashBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(hashBytes, phash.DTC(img))
	return hashBytes, nil
}

// CmpHashAndIOReader compares a hash and an io.Reader using the md5sum algorithm.
func (p *pHasher) CmpHashAndIOReader(hashA []byte, r io.Reader) error {
	hashB, err := p.GenHashFromIOReader(r)
	if err != nil {
		return err
	}

	if !bytes.Equal(hashA, hashB) {
		return ErrHashMismatch
	}
	return nil
}
