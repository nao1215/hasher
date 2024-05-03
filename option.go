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
