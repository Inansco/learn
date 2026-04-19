package main

import (
	"strconv"
	"strings"
)

func Processor(data string) string {
	words := strings.Split(data, "\n")
	for i := range words {
		words[i] = FixHex(words[i])
		words[i] = FixBin(words[i])
		words[i] = FixUp(words[i])
		words[i] = FixLow(words[i])
		words[i] = FixBin(words[i])
	}
	return strings.Join(words, "\n")
}

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

func FixBin(data string) string {
	words := strings.Fields(data)
	for i := 0; i < len(words); i++ {
		if i+1 < len(words) && words[i+1] == "(bin)" {
			bin, err := strconv.ParseInt(words[i], 2, 64)
			if err != nil {
				continue
			}
			words[i] = strconv.FormatInt(bin, 10)
			words = append(words[:i+1], words[i+2:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}

func FixUp(data string) string {
	words := strings.Fields(data)
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
		if words[i] == "(up," {
			casing := strings.Trim(words[i+1], ")")
			count, err := strconv.Atoi(casing)
			if err != nil {
				continue
			}
			for j := 0; j <= count; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}

func FixLow(data string) string {
	words := strings.Fields(data)
	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
		if words[i] == "(low," {
			casing := strings.Trim(words[i+1], ")")
			count, err := strconv.Atoi(casing)
			if err != nil {
				continue
			}
			for j := 0; j <= count; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}
