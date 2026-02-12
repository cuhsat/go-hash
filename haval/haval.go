package haval

import (
	"hash"
)

// New1283 returns a new hash.Hash computing the haval-128_3 checksum
func New1283() hash.Hash {
	h, _ := newDigest(128, 3)
	return h
}

// New1284 returns a new hash.Hash computing the haval-128_4 checksum
func New1284() hash.Hash {
	h, _ := newDigest(128, 4)
	return h
}

// New1285 returns a new hash.Hash computing the haval-128_5 checksum
func New1285() hash.Hash {
	h, _ := newDigest(128, 5)
	return h
}

// New1603 returns a new hash.Hash computing the haval-160_3 checksum
func New1603() hash.Hash {
	h, _ := newDigest(160, 3)
	return h
}

// New1604 returns a new hash.Hash computing the haval-160_4 checksum
func New1604() hash.Hash {
	h, _ := newDigest(160, 4)
	return h
}

// New1605 returns a new hash.Hash computing the haval-160_5 checksum
func New1605() hash.Hash {
	h, _ := newDigest(160, 5)
	return h
}

// New1923 returns a new hash.Hash computing the haval-192_3 checksum
func New1923() hash.Hash {
	h, _ := newDigest(192, 3)
	return h
}

// New1924 returns a new hash.Hash computing the haval-192_4 checksum
func New1924() hash.Hash {
	h, _ := newDigest(192, 4)
	return h
}

// New1925 returns a new hash.Hash computing the haval-192_5 checksum
func New1925() hash.Hash {
	h, _ := newDigest(192, 5)
	return h
}

// New2243 returns a new hash.Hash computing the haval-224_3 checksum
func New2243() hash.Hash {
	h, _ := newDigest(224, 3)
	return h
}

// New2244 returns a new hash.Hash computing the haval-224_4 checksum
func New2244() hash.Hash {
	h, _ := newDigest(224, 4)
	return h
}

// New2245 returns a new hash.Hash computing the haval-224_5 checksum
func New2245() hash.Hash {
	h, _ := newDigest(224, 5)
	return h
}

// New2563 returns a new hash.Hash computing the haval-256_3 checksum
func New2563() hash.Hash {
	h, _ := newDigest(256, 3)
	return h
}

// New2564 returns a new hash.Hash computing the haval-256_4 checksum
func New2564() hash.Hash {
	h, _ := newDigest(256, 4)
	return h
}

// New2565 returns a new hash.Hash computing the haval-256_5 checksum
func New2565() hash.Hash {
	h, _ := newDigest(256, 5)
	return h
}
