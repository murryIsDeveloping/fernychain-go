// Package hashing  is a helper package for hashing strings
package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash takes a string and returns a sha256 hash string
func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
