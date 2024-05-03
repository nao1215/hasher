package hasher

// Option sets the options for the Hasher struct.
type Option func(*Hash)

// WithUserDifinedAlgorithm is an option that sets the hash algorithm to a user-defined algorithm.
func WithUserDifinedAlgorithm(hasher Hasher) Option {
	return func(h *Hash) {
		h.hasher = hasher
	}
}

// WithMd5Sum is an option that sets the hash algorithm to MD5SUM.
func WithMd5Sum() Option {
	return func(h *Hash) {
		h.hasher = &md5sumHasher{}
	}
}

// WithSha1sum is an option that sets the hash algorithm to SHA-1.
func WithSha1sum() Option {
	return func(h *Hash) {
		h.hasher = newSHA1Hasher()
	}
}

// WithSha256sum is an option that sets the hash algorithm to SHA-256.
func WithSha256sum() Option {
	return func(h *Hash) {
		h.hasher = newSHA256Hasher()
	}
}

// WithSha512sum is an option that sets the hash algorithm to SHA-512.
func WithSha512sum() Option {
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

// WithFnv32sum is an option that sets the hash algorithm to FNV-32.
func WithFnv32sum() Option {
	return func(h *Hash) {
		h.hasher = &fnv32Hasher{}
	}
}

// WithFnv32asum is an option that sets the hash algorithm to FNV-32a.
func WithFnv32asum() Option {
	return func(h *Hash) {
		h.hasher = &fnv32aHasher{}
	}
}

// WithFnv64sum is an option that sets the hash algorithm to FNV-64.
func WithFnv64sum() Option {
	return func(h *Hash) {
		h.hasher = &fnv64Hasher{}
	}
}

// WithFnv64asum is an option that sets the hash algorithm to FNV-64a.
func WithFnv64asum() Option {
	return func(h *Hash) {
		h.hasher = &fnv64aHasher{}
	}
}

// WithFnv128sum is an option that sets the hash algorithm to FNV-128.
func WithFnv128sum() Option {
	return func(h *Hash) {
		h.hasher = newFnv128Hasher()
	}
}

// WithFnv128asum is an option that sets the hash algorithm to FNV-128a.
func WithFnv128asum() Option {
	return func(h *Hash) {
		h.hasher = newFnv128aHasher()
	}
}

// WithBlake3sum is an option that sets the hash algorithm to Blake3.
func WithBlake3sum() Option {
	return func(h *Hash) {
		h.hasher = &blake3Hasher{}
	}
}
