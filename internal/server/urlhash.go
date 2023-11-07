package server

import (
	"crypto/sha256"
	"fmt"
)

func shortHash(s string) string {
	hash32 := sha256.Sum256([]byte(s))
	strHash := fmt.Sprintf("%x", hash32)
	hash10 := []byte(strHash)[0:10]
	return string(hash10)
}

func ShortURL(longURL string) string {
	hash := shortHash(longURL)
	return fmt.Sprintf("https://urlshorter/%s", hash)
}