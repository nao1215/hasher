package hasher

// Option sets the options for the Hasher struct.
type Option func(*Hash)

// WithUserDifinedAlgorithm is an option that sets the hash algorithm to a user-defined algorithm.
func WithUserDifinedAlgorithm(hasher Hasher) Option {
	return func(h *Hash) {
		h.hasher = hasher
	}
}

// WithMd5 is an option that sets the hash algorithm to MD5SUM.
func WithMd5() Option {
	return func(h *Hash) {
		h.hasher = &md5sumHasher{}
	}
}

// WithSha1 is an option that sets the hash algorithm to SHA-1.
func WithSha1() Option {
	return func(h *Hash) {
		h.hasher = newSHA1Hasher()
	}
}

// WithSha256 is an option that sets the hash algorithm to SHA-256.
func WithSha256() Option {
	return func(h *Hash) {
		h.hasher = newSHA256Hasher()
	}
}

// WithSha512 is an option that sets the hash algorithm to SHA-512.
func WithSha512() Option {
	return func(h *Hash) {
		h.hasher = newSHA512Hasher()
	}
}

// WithPhash is an option that sets the hash algorithm to Perceptual Hash.
func WithPhash() Option {
	return func(h *Hash) {
		h.hasher = &pHasher{}
	}
}

// WithFnv32 is an option that sets the hash algorithm to FNV-32.
func WithFnv32() Option {
	return func(h *Hash) {
		h.hasher = newFnv32Hasher()
	}
}

// WithFnv32a is an option that sets the hash algorithm to FNV-32a.
func WithFnv32a() Option {
	return func(h *Hash) {
		h.hasher = newFnv32aHasher()
	}
}

// WithFnv64 is an option that sets the hash algorithm to FNV-64.
func WithFnv64() Option {
	return func(h *Hash) {
		h.hasher = newFnv64Hasher()
	}
}

// WithFnv64a is an option that sets the hash algorithm to FNV-64a.
func WithFnv64a() Option {
	return func(h *Hash) {
		h.hasher = newFnv64aHasher()
	}
}

// WithFnv128 is an option that sets the hash algorithm to FNV-128.
func WithFnv128() Option {
	return func(h *Hash) {
		h.hasher = newFnv128Hasher()
	}
}

// WithFnv128a is an option that sets the hash algorithm to FNV-128a.
func WithFnv128a() Option {
	return func(h *Hash) {
		h.hasher = newFnv128aHasher()
	}
}

// WithBlake3 is an option that sets the hash algorithm to Blake3.
func WithBlake3() Option {
	return func(h *Hash) {
		h.hasher = &blake3Hasher{}
	}
}

// WithAdler32 is an option that sets the hash algorithm to Adler-32.
func WithAdler32() Option {
	return func(h *Hash) {
		h.hasher = newAdler32Hasher()
	}
}

// WithMmh3 is an option that sets the hash algorithm to Murmur3.
func WithMmh3() Option {
	return func(h *Hash) {
		h.hasher = newMmh3Hasher()
	}
}

// WithWhirlpool is an option that sets the hash algorithm to Whirlpool.
func WithWhirlpool() Option {
	return func(h *Hash) {
		h.hasher = newWhirlpoolHasher()
	}
}

// WithCRC32 is an option that sets the hash algorithm to CRC-32.
func WithCRC32() Option {
	return func(h *Hash) {
		h.hasher = newCRC32Hasher()
	}
}

// WithXXHash is an option that sets the hash algorithm to XXHash.
func WithXXHash() Option {
	return func(h *Hash) {
		h.hasher = newXXHasher()
	}
}
