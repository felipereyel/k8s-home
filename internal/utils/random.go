package utils

import (
	"crypto/rand"
	b64 "encoding/base64"

	"github.com/google/uuid"
)

func GenerateId() string {
	return uuid.New().String()
}

func RandomString(n int) (string, error) {
	b := make([]byte, n*4)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return b64.StdEncoding.EncodeToString(b), nil
}
