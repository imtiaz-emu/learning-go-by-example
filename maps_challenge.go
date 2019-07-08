package main

import (
	"fmt"
	"strings"
)

func main() {
	wordCounter := map[string]int{}
	text := `
		Needles and pins
		Needles and pins
		sew me a sail
		To catch me the wind
	`
	words := strings.Fields(text)
	for _, word := range words {
		_, ok := wordCounter[word]
		if !ok {
			wordCounter[word] = 1
		} else {
			wordCounter[word]++
		}
	}

	fmt.Println(wordCounter)
}
