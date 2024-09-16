package utils

import (
	"math/bits"
)

func XOR(a, b []byte) []byte {
	if len(a) > len(b) {
		a = a[:len(b)]
	}
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}

func HammingDistance(a, b []byte) int {
	if len(a) != len(b) {
		panic("hammingDistance: different lengths")
	}
	var res int
	for i := range a {
		res += bits.OnesCount8(a[i] ^ b[i])
	}
	return res
}
