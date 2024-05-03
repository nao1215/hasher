// Package hasher
package hasher

import (
	"bytes"
	"encoding/hex"
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestHash_Generate(t *testing.T) {
	tests := []struct {
		name        string
		input       any
		isFile      bool
		opts        []Option
		expected    string
		expectedErr error
	}{
		{
			name:        "Generate md5sum from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{},
			expected:    "098f6bcd4621d373cade4e832627b4f6",
			expectedErr: nil,
		},
		{
			name:        "Generate md5sum from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithMd5Sum()},
			expected:    "7b4bc55c9a1295ecbd2b77a636565f27",
			expectedErr: nil,
		},
		{
			name:        "Unsupported input type",
			input:       1,
			isFile:      false,
			opts:        []Option{},
			expected:    "",
			expectedErr: ErrUnsupportedInputType,
		},
		{
			name:        "Generate user-defined hash from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithUserDifinedAlgorithm(&userHash{})},
			expected:    "74657374",
			expectedErr: nil,
		},
		{
			name:        "Generate user-defined hash from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithUserDifinedAlgorithm(&userHash{})},
			expected:    "74657374",
			expectedErr: nil,
		},
		{
			name:        "Generate sha1 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithSha1sum()},
			expected:    "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3",
			expectedErr: nil,
		},
		{
			name:        "Generate sha1 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha1sum()},
			expected:    "734202c47ed7157e3c187e1fd7ba8eb6e0b58fa4",
			expectedErr: nil,
		},
		{
			name:        "Generate sha256 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithSha256sum()},
			expected:    "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			expectedErr: nil,
		},
		{
			name:        "Generate sha256 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha256sum()},
			expected:    "f2e0d62d29081f468bf7ac70415fc4cc391d877fb7ef09c614c7bdd5f9175b99",
			expectedErr: nil,
		},
		{
			name:        "Generate sha512 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithSha512sum()},
			expected:    "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff",
			expectedErr: nil,
		},
		{
			name:        "Generate sha512 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha512sum()},
			expected:    "9e7021341882d2a4cae911cf08b0312a10c8edff7aa279adb43b2c2646bece9281da78e2d6e84c048b9ff70730990bfd201240c18b6e053b2027605690671418",
			expectedErr: nil,
		},
		{
			name:        "Failed to generate perceptual hash from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithPhash()},
			expected:    "",
			expectedErr: ErrPhashNotSupportedString,
		},
		{
			name:        "Generate perceptual hash from io.Reader",
			input:       filepath.Join("testdata", "test.jpg"),
			isFile:      true,
			opts:        []Option{WithPhash()},
			expected:    "6917092734e3ec3a",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := NewHash(tt.opts...)

			var input any
			input = tt.input
			if tt.isFile {
				i, ok := tt.input.(string)
				if !ok {
					t.Fatalf("input is not a string")
				}

				f, err := os.Open(filepath.Clean(i))
				if err != nil {
					t.Fatalf("os.Open() error = %v", err)
				}
				defer f.Close() //nolint:errcheck

				input = f
			}

			got, err := h.Generate(input)
			if tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Hash.Generate() error = %v, want %v", err, tt.expectedErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("Hash.Generate() error = %v", err)
			}

			expectedBytes, err := hex.DecodeString(tt.expected)
			if err != nil {
				t.Fatalf("hex.DecodeString() error = %v", err)
			}

			if !bytes.Equal(got, expectedBytes) {
				t.Errorf("Hash.Generate() = %x, want %x", got, expectedBytes)
			}
		})
	}
}
func TestHash_Compare(t *testing.T) {
	tests := []struct {
		name        string
		hash        string
		input       interface{}
		isFile      bool
		opts        []Option
		expectedErr error
	}{
		{
			name:        "Compare md5sum hash and string",
			hash:        "098f6bcd4621d373cade4e832627b4f6",
			input:       "test",
			isFile:      false,
			opts:        []Option{},
			expectedErr: nil,
		},
		{
			name:        "Compare md5sum hash and io.Reader",
			hash:        "7b4bc55c9a1295ecbd2b77a636565f27",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithMd5Sum()},
			expectedErr: nil,
		},
		{
			name:        "Hash mismatch: input type is string",
			hash:        "098f6bcd4621d373cade4e832627b4f6",
			input:       "mismatch_string",
			isFile:      false,
			opts:        []Option{},
			expectedErr: ErrHashMismatch,
		},
		{
			name:        "Hash mismatch: input type is io.Reader",
			hash:        "7b4bc55c9a1295ecbd2b77a636565f27",
			input:       filepath.Join("testdata", "mismatch.txt"),
			isFile:      true,
			opts:        []Option{WithMd5Sum()},
			expectedErr: ErrHashMismatch,
		},
		{
			name:        "Unsupported input type",
			hash:        "098f6bcd4621d373cade4e832627b4f6",
			input:       1,
			isFile:      false,
			opts:        []Option{},
			expectedErr: ErrUnsupportedInputType,
		},
		{
			name:        "Compare user-defined hash and string",
			hash:        "0c",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithUserDifinedAlgorithm(&userHash{})},
			expectedErr: nil,
		},
		{
			name:        "Compare user-defined hash and io.Reader",
			hash:        "22",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithUserDifinedAlgorithm(&userHash{})},
			expectedErr: nil,
		},
		{
			name:        "Compare sha1 hash and string",
			hash:        "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithSha1sum()},
			expectedErr: nil,
		},
		{
			name:        "Compare sha1 hash and io.Reader",
			hash:        "734202c47ed7157e3c187e1fd7ba8eb6e0b58fa4",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha1sum()},
			expectedErr: nil,
		},
		{
			name:        "Compare sha256 hash and string",
			hash:        "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithSha256sum()},
			expectedErr: nil,
		},
		{
			name:        "Compare s sha256 hash and io.Reader",
			hash:        "f2e0d62d29081f468bf7ac70415fc4cc391d877fb7ef09c614c7bdd5f9175b99",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha256sum()},
			expectedErr: nil,
		},
		{
			name:        "Compare sha512 hash and string",
			hash:        "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithSha512sum()},
			expectedErr: nil,
		},
		{
			name:        "Compare sha512 hash and io.Reader",
			hash:        "9e7021341882d2a4cae911cf08b0312a10c8edff7aa279adb43b2c2646bece9281da78e2d6e84c048b9ff70730990bfd201240c18b6e053b2027605690671418",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha512sum()},
			expectedErr: nil,
		},
		{
			name:        "Failed to compare perceptual hash and string",
			hash:        "6917092734e3ec3a",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithPhash()},
			expectedErr: ErrPhashNotSupportedString,
		},
		{
			name:        "Compare perceptual hash and io.Reader",
			hash:        "6917092734e3ec3a",
			input:       filepath.Join("testdata", "test.jpg"),
			isFile:      true,
			opts:        []Option{WithPhash()},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			hashBytes, err := hex.DecodeString(tt.hash)
			if err != nil {
				t.Fatalf("hex.DecodeString() error = %v", err)
			}

			h := NewHash(tt.opts...)

			var input any
			input = tt.input
			if tt.isFile {
				if _, ok := tt.input.(string); ok {
					i, ok := tt.input.(string)
					if !ok {
						t.Fatalf("input is not a string")
					}

					f, err := os.Open(filepath.Clean(i))
					if err != nil {
						t.Fatalf("os.Open() error = %v", err)
					}
					defer f.Close() //nolint:errcheck

					input = f
				}
			}

			err = h.Compare(hashBytes, input)
			if tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Errorf("Hash.Compare() error = %v, want %v", err, tt.expectedErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("Hash.Compare() error = %v", err)
			}
		})
	}
}

type userHash struct{}

func (u *userHash) GenHashFromString(_ string) ([]byte, error) {
	return []byte{116, 101, 115, 116}, nil // "test"
}

func (u *userHash) GenHashFromIOReader(_ io.Reader) ([]byte, error) {
	return []byte{116, 101, 115, 116}, nil // "test"
}

func (u *userHash) CmpHashAndString(hash []byte, s string) error {
	return nil
}

func (u *userHash) CmpHashAndIOReader(hash []byte, r io.Reader) error {
	return nil
}
