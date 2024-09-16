package stream

import (
	"bytes"
	"github.com/aituglo/cryptack/pkg/utils"
	"os"
	"strings"
	"testing"
)

func TestFindSingleByteXORKey(t *testing.T) {
	corpus := utils.GetLanguageFrequency("english")
	res, _, _ := FindSingleByteXORKey(utils.DecodeHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"), corpus)
	t.Logf("%s", res)
}

func TestDetectSingleByteXORKey(t *testing.T) {
	text, err := os.ReadFile("testfiles/single_char_xor.txt")
	if err != nil {
		t.Fatal("failed to read file:", err)
	}
	var bestScore float64
	var res []byte
	corpus := utils.GetLanguageFrequency("english")
	for _, line := range strings.Split(string(text), "\n") {
		out, _, score := FindSingleByteXORKey(utils.DecodeHex(line), corpus)
		if score > bestScore {
			res = out
			bestScore = score
		}
	}
	t.Logf("%s", res)
}

func TestRepeatingKeyXOR(t *testing.T) {
	input := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	res := RepeatingKeyXOR(input, []byte("ICE"))
	if !bytes.Equal(res, utils.DecodeHex("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")) {
		t.Error("wrong result:", res)
	}
}
