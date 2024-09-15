package utils

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
