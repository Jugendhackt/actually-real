package token

import (
	"crypto/rand"
	"encoding/base64"
)

func generateRandom(length int) string {
	b := make([]byte, length)

	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(b)
}
