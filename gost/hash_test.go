package gost

import (
	"fmt"
	"hash"
	"testing"
)

func TestHashInterface(t *testing.T) {
	h := New()
	var _ hash.Hash = h
}

func TestVectors(t *testing.T) {
	h := New()
	_, _ = h.Write([]byte("a"))
	hashed := h.Sum(nil)

	if len(hashed) == 0 {
		t.Error("Hash error")
	}
}

func Test_Check(t *testing.T) {
	in := []byte("nonce-asdfg")
	check := "08246810594b02216e6f6e633d2d39733c666352565b5b6a344e580c0b4e580c"

	h := New()
	_, _ = h.Write(in)

	out := h.Sum(nil)

	if fmt.Sprintf("%x", out) != check {
		t.Errorf("Check error. got %x, want %s", out, check)
	}
}
