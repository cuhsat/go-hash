package haval

import (
	"hash"
)

// New128_3 returns a new hash.Hash computing the haval-128_3 checksum
func New128_3() hash.Hash {
	h, _ := newDigest(128, 3)
	return h
}

// New128_4 returns a new hash.Hash computing the haval-128_4 checksum
func New128_4() hash.Hash {
	h, _ := newDigest(128, 4)
	return h
}

// New128_5 returns a new hash.Hash computing the haval-128_5 checksum
func New128_5() hash.Hash {
	h, _ := newDigest(128, 5)
	return h
}

// New160_3 returns a new hash.Hash computing the haval-160_3 checksum
func New160_3() hash.Hash {
	h, _ := newDigest(160, 3)
	return h
}

// New160_4 returns a new hash.Hash computing the haval-160_4 checksum
func New160_4() hash.Hash {
	h, _ := newDigest(160, 4)
	return h
}

// New160_5 returns a new hash.Hash computing the haval-160_5 checksum
func New160_5() hash.Hash {
	h, _ := newDigest(160, 5)
	return h
}

// New192_3 returns a new hash.Hash computing the haval-192_3 checksum
func New192_3() hash.Hash {
	h, _ := newDigest(192, 3)
	return h
}

// New192_4 returns a new hash.Hash computing the haval-192_4 checksum
func New192_4() hash.Hash {
	h, _ := newDigest(192, 4)
	return h
}

// New192_5 returns a new hash.Hash computing the haval-192_5 checksum
func New192_5() hash.Hash {
	h, _ := newDigest(192, 5)
	return h
}

// New224_3 returns a new hash.Hash computing the haval-224_3 checksum
func New224_3() hash.Hash {
	h, _ := newDigest(224, 3)
	return h
}

// New224_4 returns a new hash.Hash computing the haval-224_4 checksum
func New224_4() hash.Hash {
	h, _ := newDigest(224, 4)
	return h
}

// New224_5 returns a new hash.Hash computing the haval-224_5 checksum
func New224_5() hash.Hash {
	h, _ := newDigest(224, 5)
	return h
}

// New256_3 returns a new hash.Hash computing the haval-256_3 checksum
func New256_3() hash.Hash {
	h, _ := newDigest(256, 3)
	return h
}

// New256_4 returns a new hash.Hash computing the haval-256_4 checksum
func New256_4() hash.Hash {
	h, _ := newDigest(256, 4)
	return h
}

// New256_5 returns a new hash.Hash computing the haval-256_5 checksum
func New256_5() hash.Hash {
	h, _ := newDigest(256, 5)
	return h
}
