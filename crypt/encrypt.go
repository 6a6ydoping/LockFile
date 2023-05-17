package crypt

import (
	"encoding/base64"
)

func EncryptBytes(bs []byte, secretKey string) (string, error) {
	return "", nil
}

// AES ENCRYPTION

func Encode(bs []byte) string {
	return base64.StdEncoding.EncodeToString(bs)
}
