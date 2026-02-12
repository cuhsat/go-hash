// Package skein implements the Skein-512 hash function, MAC, and stream cipher
// as defined in "The Skein Hash Function Family, v1.3".
package skein

import (
	"hash"
)

// NewHash returns hash.Hash calculating checksum of the given length in bytes
// (for example, to calculate 256-bit hash, outLen must be set to 32).
func NewHash(outLen uint64) hash.Hash {
	return hash.Hash(New(outLen, nil))
}

// NewHash224 returns a new hash.Hash computing the Skein-224 checksum
func NewHash224() hash.Hash {
	return NewHash(28)
}

// NewHash256 returns a new hash.Hash computing the Skein-256 checksum
func NewHash256() hash.Hash {
	return NewHash(32)
}

// NewHash384 returns a new hash.Hash computing the Skein-384 checksum
func NewHash384() hash.Hash {
	return NewHash(48)
}

// NewHash512 returns a new hash.Hash computing the Skein-512 checksum
func NewHash512() hash.Hash {
	return NewHash(64)
}
