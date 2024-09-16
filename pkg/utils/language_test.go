package utils

import "testing"

func TestFrequencyAnalysis(t *testing.T) {
	res := FrequencyAnalysis("hello")
	if res['h'] != 0.2 {
		t.Fatal("wrong frequency")
	}
	if res['e'] != 0.2 {
		t.Fatal("wrong frequency")
	}
	if res['l'] != 0.4 {
		t.Fatal("wrong frequency")
	}
	if res['o'] != 0.2 {
		t.Fatal("wrong frequency")
	}
}
