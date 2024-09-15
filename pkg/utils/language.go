package utils

import (
	"unicode/utf8"
)

func FrequencyAnalysis(text string) map[rune]float64 {
	res := make(map[rune]float64)
	for _, c := range text {
		res[c]++
	}
	total := utf8.RuneCountInString(text)
	for char := range res {
		res[char] = res[char] / float64(total)
	}
	return res
}

func ScoreText(text string, c map[rune]float64) float64 {
	var score float64
	for _, char := range text {
		score += c[char]
	}
	return score / float64(utf8.RuneCountInString(text))
}
