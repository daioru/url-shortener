package service

import (
	"crypto/rand"
	"encoding/base64"
)

func generateShortURL() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:6]
}
