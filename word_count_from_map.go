package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
		Needles and pins
		Needles and pins
		Sew me a sail
		To catch me the wind
	`
	//empty map
	wordCountMap := map[string]int{}
	words := strings.Fields(text)

	for _,word := range words {
		wordCountMap[strings.ToLower(word)]++ 
	}

	fmt.Println(wordCountMap)

}