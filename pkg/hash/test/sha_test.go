// Package hash
package test

import (
	"github.com/dienggo/diego/pkg/hash"
	"testing"
)

func TestSHA256(t *testing.T) {
	hashed := hash.SHA256("test-3")
	expected := "2a8c9f051e91be1d0f801980a9e87f8495582668d966b633bfde5d8a93d0e049"
	if hashed != expected {
		t.Fatalf("expected %s, actual %s", expected, hashed)
	}
	t.Logf("expected %s, actual %s", expected, hashed)
}

func TestSHA1(t *testing.T) {
	hashed := hash.SHA1("test-2")
	expected := "987d4cc40e5e5c0fdb68adc87e3445740670e2d0"
	if hashed != expected {
		t.Fatalf("expected %s, actual %s", expected, hashed)
	}
	t.Logf("expected %s, actual %s", expected, hashed)
}
