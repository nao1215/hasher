package hasher

import (
	"hash/fnv"
)

// newFnv128Hasher creates a new Hasher instance for FNV-128 algorithm.
func newFnv128Hasher() Hasher {
	return &hasher{HashFunc: fnv.New128}
}

// newFnv128aHasher creates a new Hasher instance for FNV-128a algorithm.
func newFnv128aHasher() Hasher {
	return &hasher{HashFunc: fnv.New128a}
}

// newFnv32Hasher creates a new Hasher instance for FNV-32 algorithm.
func newFnv32Hasher() Hasher {
	return &hasher32{HashFunc: fnv.New32}
}

// newFnv32aHasher creates a new Hasher instance for FNV-32a algorithm.
func newFnv32aHasher() Hasher {
	return &hasher32{HashFunc: fnv.New32a}
}

// newFnv64Hasher creates a new Hasher instance for FNV-64 algorithm.
func newFnv64Hasher() Hasher {
	return &hasher64{HashFunc: fnv.New64}
}

// newFnv64aHasher creates a new Hasher instance for FNV-64a algorithm.
func newFnv64aHasher() Hasher {
	return &hasher64{HashFunc: fnv.New64a}
}
