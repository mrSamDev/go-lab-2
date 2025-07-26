package main

import (
	"fmt"
	"strings"
)

const book = `In the beginning God created the heavens and the earth. Now the earth was formless and empty, darkness was over the surface of the deep, and the Spirit of God was hovering over the waters. And God said, "Let there be light," and there was light.`

func main() {

	words := strings.Fields(book)

	var wordFrequency = make(map[string]int)

	for _, word := range words {
		word = strings.ToLower(word)
		wordFrequency[word]++
	}

	fmt.Println("Word Frequency:", wordFrequency)

}
