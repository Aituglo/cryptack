package utils

import (
	"embed"
	"log"
	"unicode/utf8"
)

//go:embed corpus/*
var corpus embed.FS

func corpusFromFile(name string) map[rune]float64 {
	text, err := corpus.ReadFile(name)
	if err != nil {
		log.Fatal("failed to read corpus file:", err)
	}
	return FrequencyAnalysis(string(text))
}

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

func GetLanguageFrequency(language string) map[rune]float64 {
	var frequency map[rune]float64

	switch language {
	case "english":
		frequency = corpusFromFile("corpus/english.txt")
	default:
		frequency = corpusFromFile("corpus/english.txt")
	}

	return frequency

}

func ScoreText(text string, frequency map[rune]float64) float64 {
	var score float64

	for _, char := range text {
		score += frequency[char]
	}
	return score / float64(utf8.RuneCountInString(text))
}
