package main

import (
	"fmt"
	"strings"
)

func WordCount(str string) map[string]int {
	words := strings.Fields(str)
	words_len := len(words)

	words_map := make(map[string]int)

	// update the map
	for i := 0; i < words_len; i++ {
		words_map[words[i]] += 1
	}

	return words_map
}

func main() {
	_map := WordCount("Hello Hello World World World World!")
	fmt.Println(_map["Hello"])       // 2
	fmt.Println(_map["World"])       // 3
	fmt.Println(_map["Woqwewqerld"]) // 0
}
