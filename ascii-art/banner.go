package main

import (
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	characters := make(map[rune][]string)
	for i := 0; i < len(lines); i +=9 {
		char := rune(32 + i/9)
		artLines := lines[i+1 : i+9]
		characters[char] = artLines
	}
	return characters, nil
}