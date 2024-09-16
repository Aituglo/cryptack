package utils

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func DecodeHex(s string) []byte {
	v, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal("failed to decode hex:", s)
	}
	return v
}

func EncodeHex(b []byte) string {
	return hex.EncodeToString(b)
}

func HexToBase64(input string) (string, error) {
	value, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(value), nil
}

func Base64ToHex(input string) (string, error) {
	value, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(value), nil
}
