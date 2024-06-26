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
			opts:        []Option{WithMd5()},
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
			opts:        []Option{WithSha1()},
			expected:    "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3",
			expectedErr: nil,
		},
		{
			name:        "Generate sha1 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha1()},
			expected:    "734202c47ed7157e3c187e1fd7ba8eb6e0b58fa4",
			expectedErr: nil,
		},
		{
			name:        "Generate sha256 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithSha256()},
			expected:    "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			expectedErr: nil,
		},
		{
			name:        "Generate sha256 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha256()},
			expected:    "f2e0d62d29081f468bf7ac70415fc4cc391d877fb7ef09c614c7bdd5f9175b99",
			expectedErr: nil,
		},
		{
			name:        "Generate sha512 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithSha512()},
			expected:    "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff",
			expectedErr: nil,
		},
		{
			name:        "Generate sha512 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha512()},
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
		{
			name:        "Generate fnv32 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv32()},
			expected:    "bc2c0be9",
			expectedErr: nil,
		},
		{
			name:        "Generate fnv32 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv32()},
			expected:    "43ada300",
			expectedErr: nil,
		},
		{
			name:        "Generate fnv32a from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv32a()},
			expected:    "afd071e5",
			expectedErr: nil,
		},
		{
			name:        "Generate fnv32a from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv32a()},
			expected:    "2358b7f6",
			expectedErr: nil,
		},
		{
			name:        "Generate fnv64 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv64()},
			expected:    "8c093f7e9fccbf69",
			expectedErr: nil,
		},
		{
			name:        "Generate fnv64 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv64()},
			expected:    "137ece8257075460",
			expectedErr: nil,
		},
		{
			name:        "Generate fnv64a from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv64a()},
			expected:    "f9e6e6ef197c2b25",
			expectedErr: nil,
		},
		{
			name:        "Generate fnv64a from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv64a()},
			expected:    "a696a78da2a53bf6",
			expectedErr: nil,
		},
		{
			name:        "Generate fnv128 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv128()},
			expected:    "66ab2a8b6f757277b806e89c56faf339",
			expectedErr: nil,
		},
		{
			name:        "Generate fnv128 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv128()},
			expected:    "68958d70aeb8df5cd163c8a14d117b08",
			expectedErr: nil,
		},
		{
			name:        "Generate fnv128a from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv128a()},
			expected:    "69d061a9c5757277b806e99413dd99a5",
			expectedErr: nil,
		},
		{
			name:        "Generate fnv128a from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv128a()},
			expected:    "2736af9add0bf02253c1651da64b19a6",
			expectedErr: nil,
		},
		{
			name:        "Generate blake3 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithBlake3()},
			expected:    "4878ca0425c739fa427f7eda20fe845f6b2e46ba5fe2a14df5b1e32f50603215c82f77a5bd07f7048a95a699e056d0e32bd2bdadc37ee096719c3d9ec12f29a6",
			expectedErr: nil,
		},
		{
			name:        "Generate blake3 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithBlake3()},
			expected:    "a1b692dcd30210a76e8d24cb593a0b2f8d307dbe4dfdefd7053f725d6b996fe9ee2ea450c9c9afa2f654640ec4cf113c3420d2f3b71edd7c55ca11d918d1af2b",
			expectedErr: nil,
		},
		{
			name:        "Generate adler32 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithAdler32()},
			expected:    "045d01c1",
			expectedErr: nil,
		},
		{
			name:        "Generate adler32 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithAdler32()},
			expected:    "0368237e",
			expectedErr: nil,
		},
		{
			name:        "Generate mmh3 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithMmh3()},
			expected:    "9de1bd74cc287dac824dbdf93182129a",
			expectedErr: nil,
		},
		{
			name:        "Generate mmh3 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithMmh3()},
			expected:    "44419d48e3c263f195119454fde57819",
			expectedErr: nil,
		},
		{
			name:        "Generate whirlpool from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithWhirlpool()},
			expected:    "b913d5bbb8e461c2c5961cbe0edcdadfd29f068225ceb37da6defcf89849368f8c6c2eb6a4c4ac75775d032a0ecfdfe8550573062b653fe92fc7b8fb3b7be8d6",
			expectedErr: nil,
		},
		{
			name:        "Generate whirlpool from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithWhirlpool()},
			expected:    "00a1616b2be44c0d92fef873630421e01fd2d6c346502281595b818ca262f3211e52d72adc3cdc396c2953ba8e62d54805f0ae791e8d032c3df3b4ab84c2a4f3",
			expectedErr: nil,
		},
		{
			name:        "Generate crc32 from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithCRC32()},
			expected:    "d87f7e0c",
			expectedErr: nil,
		},
		{
			name:        "Generate crc32 from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithCRC32()},
			expected:    "5c98c4e4",
			expectedErr: nil,
		},
		{
			name:        "Generate xxHash from string",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithXXHash()},
			expected:    "4fdcca5ddb678139",
			expectedErr: nil,
		},
		{
			name:        "Generate xxHash from io.Reader",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithXXHash()},
			expected:    "29081d4d3fb56bc6",
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
			opts:        []Option{WithMd5()},
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
			opts:        []Option{WithMd5()},
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
			opts:        []Option{WithSha1()},
			expectedErr: nil,
		},
		{
			name:        "Compare sha1 hash and io.Reader",
			hash:        "734202c47ed7157e3c187e1fd7ba8eb6e0b58fa4",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha1()},
			expectedErr: nil,
		},
		{
			name:        "Compare sha256 hash and string",
			hash:        "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithSha256()},
			expectedErr: nil,
		},
		{
			name:        "Compare s sha256 hash and io.Reader",
			hash:        "f2e0d62d29081f468bf7ac70415fc4cc391d877fb7ef09c614c7bdd5f9175b99",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha256()},
			expectedErr: nil,
		},
		{
			name:        "Compare sha512 hash and string",
			hash:        "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithSha512()},
			expectedErr: nil,
		},
		{
			name:        "Compare sha512 hash and io.Reader",
			hash:        "9e7021341882d2a4cae911cf08b0312a10c8edff7aa279adb43b2c2646bece9281da78e2d6e84c048b9ff70730990bfd201240c18b6e053b2027605690671418",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithSha512()},
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
		{
			name:        "Compare fnv32 hash and string",
			hash:        "bc2c0be9",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv32()},
			expectedErr: nil,
		},
		{
			name:        "Compare fnv32 hash and io.Reader",
			hash:        "43ada300",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv32()},
			expectedErr: nil,
		},
		{
			name:        "Compare fnv32a hash and string",
			hash:        "afd071e5",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv32a()},
			expectedErr: nil,
		},
		{
			name:        "Compare fnv32a hash and io.Reader",
			hash:        "2358b7f6",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv32a()},
			expectedErr: nil,
		},
		{
			name:        "Compare fnv64 hash and string",
			hash:        "8c093f7e9fccbf69",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv64()},
			expectedErr: nil,
		},
		{
			name:        "Compare fnv64 hash and io.Reader",
			hash:        "137ece8257075460",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv64()},
			expectedErr: nil,
		},
		{
			name:        "Compare fnv64a hash and string",
			hash:        "f9e6e6ef197c2b25",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv64a()},
			expectedErr: nil,
		},
		{
			name:        "Compare fnv64a hash and io.Reader",
			hash:        "a696a78da2a53bf6",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv64a()},
			expectedErr: nil,
		},
		{
			name:        "Compare fnv128 hash and string",
			hash:        "66ab2a8b6f757277b806e89c56faf339",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv128()},
			expectedErr: nil,
		},
		{
			name:        "Compare fnv128 hash and io.Reader",
			hash:        "68958d70aeb8df5cd163c8a14d117b08",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv128()},
			expectedErr: nil,
		},
		{
			name:        "Compare fnv128a hash and string",
			hash:        "69d061a9c5757277b806e99413dd99a5",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithFnv128a()},
			expectedErr: nil,
		},
		{
			name:        "Compare fnv128a hash and io.Reader",
			hash:        "2736af9add0bf02253c1651da64b19a6",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithFnv128a()},
			expectedErr: nil,
		},
		{
			name:        "Compare blake3 hash and string",
			hash:        "4878ca0425c739fa427f7eda20fe845f6b2e46ba5fe2a14df5b1e32f50603215c82f77a5bd07f7048a95a699e056d0e32bd2bdadc37ee096719c3d9ec12f29a6",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithBlake3()},
			expectedErr: nil,
		},
		{
			name:        "Compare blake3 hash and io.Reader",
			hash:        "a1b692dcd30210a76e8d24cb593a0b2f8d307dbe4dfdefd7053f725d6b996fe9ee2ea450c9c9afa2f654640ec4cf113c3420d2f3b71edd7c55ca11d918d1af2b",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithBlake3()},
			expectedErr: nil,
		},
		{
			name:        "Compare adler32 hash and string",
			hash:        "045d01c1",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithAdler32()},
			expectedErr: nil,
		},
		{
			name:        "Compare adler32 hash and io.Reader",
			hash:        "0368237e",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithAdler32()},
			expectedErr: nil,
		},
		{
			name:        "Compare mmh3 hash and string",
			hash:        "9de1bd74cc287dac824dbdf93182129a",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithMmh3()},
			expectedErr: nil,
		},
		{
			name:        "Compare mmh3 hash and io.Reader",
			hash:        "44419d48e3c263f195119454fde57819",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithMmh3()},
			expectedErr: nil,
		},
		{
			name:        "Compare whirlpool hash and string",
			hash:        "b913d5bbb8e461c2c5961cbe0edcdadfd29f068225ceb37da6defcf89849368f8c6c2eb6a4c4ac75775d032a0ecfdfe8550573062b653fe92fc7b8fb3b7be8d6",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithWhirlpool()},
			expectedErr: nil,
		},
		{
			name:        "Compare whirlpool hash and io.Reader",
			hash:        "00a1616b2be44c0d92fef873630421e01fd2d6c346502281595b818ca262f3211e52d72adc3cdc396c2953ba8e62d54805f0ae791e8d032c3df3b4ab84c2a4f3",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithWhirlpool()},
			expectedErr: nil,
		},
		{
			name:        "Compare crc32 hash and string",
			hash:        "d87f7e0c",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithCRC32()},
			expectedErr: nil,
		},
		{
			name:        "Compare crc32 hash and io.Reader",
			hash:        "5c98c4e4",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithCRC32()},
			expectedErr: nil,
		},
		{
			name:        "Compare xxHash hash and string",
			hash:        "4fdcca5ddb678139",
			input:       "test",
			isFile:      false,
			opts:        []Option{WithXXHash()},
			expectedErr: nil,
		},
		{
			name:        "Compare xxHash hash and io.Reader",
			hash:        "29081d4d3fb56bc6",
			input:       filepath.Join("testdata", "test.txt"),
			isFile:      true,
			opts:        []Option{WithXXHash()},
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
