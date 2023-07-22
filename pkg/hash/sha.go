// Package hash
package hash

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// SHA1 hasher
func SHA1(msg string) string {
	h := sha1.New()
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256 hasher
func SHA256(msg string) string {
	h := sha256.New()
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}
