package stream

import (
	"github.com/aituglo/cryptack/pkg/utils"
	"math"
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

func FindRepeatingXORSize(in []byte, max int) int {
	var res int
	bestScore := math.MaxFloat64
	for keyLen := 2; keyLen < max; keyLen++ {
		a, b := in[:keyLen*4], in[keyLen*4:keyLen*4*2]
		score := float64(utils.HammingDistance(a, b)) / float64(keyLen)
		if score < bestScore {
			res = keyLen
			bestScore = score
		}
	}
	return res
}

func FindRepeatingXORKey(in []byte, frequency map[rune]float64, max int) []byte {
	keySize := FindRepeatingXORSize(in, max)
	column := make([]byte, (len(in)+keySize-1)/keySize)
	key := make([]byte, keySize)
	for col := 0; col < keySize; col++ {
		for row := range column {
			if row*keySize+col >= len(in) {
				continue
			}
			column[row] = in[row*keySize+col]
		}
		_, k, _ := FindSingleByteXORKey(column, frequency)
		key[col] = k
	}
	return key
}
