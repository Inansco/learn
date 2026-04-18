package main

import (
	"strconv"
	"strings"
)

// func Processor(data string) string {
// 	words := strings.Split(data, ",")
// 	for i := range words {
// 		words[i] = FixHex(data)
// 		return data
// 	}
// 	return strings.Join(data, " ")
// }

func FixHex(data string) string {
	words := strings.Fields(data)
	for i := 0; i < len(words); i++ {
		if i+1 < len(words) && words[i+1] == "(hex)" {
			hex, err := strconv.ParseInt(words[i], 16, 64)
			if err != nil {
				continue
			}
			words[i] = strconv.FormatInt(hex, 10)
			words = append(words[:i+1], words[i+2:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}

