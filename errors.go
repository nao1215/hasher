package hasher

import "errors"

var (
	// ErrUnsupportedInputType is an error that is returned when the input type is not supported.
	ErrUnsupportedInputType = errors.New("unsupported input type")
	// ErrHashMismatch is an error that is returned when the hash and the input do not match.
	ErrHashMismatch = errors.New("hash mismatch")
	// ErrPhashNotSupportedString is an error that is returned when phash does not support string input.
	ErrPhashNotSupportedString = errors.New("phash does not support string input")
)
