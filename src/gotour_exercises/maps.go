package main

import (
	"strings"
)

func wordCount(str string) map[string]int {
	words := strings.Fields(str)
	words_len := len(words)

	words_map := make(map[string]int)

	// zero the map
	for i := 0; i < words_len; i++ {
		words_map[words[i]] = 0
	}

	// update the map
	for i := 0; i < words_len; i++ {
		words_map[words[i]] += 1
	}

	return map[string]int{"Hi": 1}
}
