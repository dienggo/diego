package hash

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash create md5 digest from string
func MD5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
