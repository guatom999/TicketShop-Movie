package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomString() string {

	byteString := make([]byte, 4)

	_, err := rand.Read(byteString)
	if err != nil {
		return ""
	}

	randomString := hex.EncodeToString(byteString)

	return randomString[:7]

}
