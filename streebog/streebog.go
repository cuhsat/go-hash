package streebog

import "hash"

// New returns a new hash.Hash computing the Streebog checksum
func New(hashsize int) (hash.Hash, error) {
	return newDigest(hashsize)
}

// New256 returns a new hash.Hash computing the Streebog checksum
func New256() hash.Hash {
	h, _ := New(256)
	return h
}

// New512 returns a new hash.Hash computing the Streebog checksum
func New512() hash.Hash {
	h, _ := New(512)
	return h
}
