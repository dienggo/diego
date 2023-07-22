// Package hash
package test

import (
	"bytes"
	"github.com/daewu14/golang-base/pkg/hash"
	"testing"
)

func TestHmac256(t *testing.T) {

	want := "0329a06b62cd16b33eb6792be8c60b158d89a2ee3a876fce9a881ebb488c0914"
	if result := hash.Hmac256("test", "secret"); result == want {
		t.Logf("expected: %v, got: %v", want, result)
	} else {
		t.Fatalf("expected: %v, got: %v", want, result)

	}
}

func TestHmacComparator(t *testing.T) {
	encrypted := "0329a06b62cd16b33eb6792be8c60b158d89a2ee3a876fce9a881ebb488c0914"
	want := true
	if result := hash.HmacComparator("test", encrypted, "secret"); result == want {
		t.Logf("expected: %v, got: %v", want, result)
	} else {
		t.Fatalf("expected: %v, got: %v", want, result)

	}
}

func TestHmac256Raw(t *testing.T) {

	expected := []byte{
		136, 205, 33, 8, 181, 52, 125, 151, 60, 243, 156, 223, 144, 83, 215, 221, 66, 112, 72, 118, 216, 201, 169, 189, 142, 45, 22, 130, 89, 211, 221, 247,
	}
	b := hash.Hmac256Raw("test", "test")

	if !bytes.Equal(expected, b) {
		t.Fatalf("expected: %v, got: %v", expected, b)
	}

	t.Logf("expected: %v, got: %v", expected, b)

}
