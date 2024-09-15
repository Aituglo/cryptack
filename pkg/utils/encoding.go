package utils

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(input string) (string, error) {
	value, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(value), nil
}
