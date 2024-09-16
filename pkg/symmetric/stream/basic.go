package stream

import (
	"github.com/aituglo/cryptack/pkg/utils"
)

// Encrypt
func SingleByteXOR(input []byte, key byte) []byte {
	res := make([]byte, len(input))
	for i := range input {
		res[i] = input[i] ^ key
	}
	return res
}

func RepeatingKeyXOR(input, key []byte) []byte {
	res := make([]byte, len(input))
	for i := range input {
		res[i] = input[i] ^ key[i%len(key)]
	}
	return res
}

// Attack
func FindSingleByteXORKey(input []byte, frequency map[rune]float64) (res []byte, key byte, score float64) {
	for k := 0; k < 256; k++ {
		out := SingleByteXOR(input, byte(k))
		s := utils.ScoreText(string(out), frequency)
		if s > score {
			res = out
			score = s
			key = byte(k)
		}
	}
	return
}
