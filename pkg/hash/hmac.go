// Package hash
package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Hmac256 generate hash hmac sha256
func Hmac256(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

// HmacComparator compare signature request body with hmac hash from payload
func HmacComparator(message string, messageHmac string, secret string) bool {
	return messageHmac == Hmac256(message, secret)
}

func Hmac256Raw(src, secret string) []byte {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	return h.Sum(nil)
}