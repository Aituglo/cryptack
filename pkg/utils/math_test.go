package utils

import (
	"bytes"
	"testing"
)

func TestXOR(t *testing.T) {
	res := XOR(DecodeHex("1c0111001f010100061a024b53535009181c"),
		DecodeHex("686974207468652062756c6c277320657965"))
	if !bytes.Equal(res, DecodeHex("746865206b696420646f6e277420706c6179")) {
		t.Errorf("wrong result: %x", res)
	}
}

func TestHammingDistance(t *testing.T) {
	res := HammingDistance([]byte("this is a test"), []byte("wokka wokka!!!"))
	if res != 37 {
		t.Fatal("wrong Hamming distance:", res)
	}
}
