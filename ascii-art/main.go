package main

import (
	"strings"
)

func asciiArt(word string, characters []string) string {
	if word == "" {
		return ""
	}
	result := ""
	words := SplitInput(word)
	for _, w := range words {
		var allLines [][]string
		for _, letter := range w {
			position := int(letter) - 32
			lines := strings.Split(characters[position], "\n")
			allLines = append(allLines, lines)
		}
		for lineNum := 0; lineNum < 8; lineNum++ {
			for _, lines := range allLines {
				result += lines[lineNum]
			}
			result += "\n"
		}

	}
	return result
}
