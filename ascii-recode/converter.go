package main

import "strings"

func StringToArt(input string) string {
	numbers := map[rune][5]string{
		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},

		'9': {
			" ___ ",
			"|___|",
			"|___|",
			" ___|",
			"     ",
		},
	}
	if input == "" {
		return ""
	}
	segments := strings.Split(input, "\n")
	var segmentsEnd []string
	for _, segment := range segments {
		for _, char := range segment {
			if char < '1' || char > '9' {
				return ""
			}
		}
		var rows [5]string
		for _, char := range segment {
			numbers, ok := numbers[char]
			if !ok {
				
			}
			for i := 0; i < 5; i++ {
				rows[i] += numbers[i]
			}
		}
		segmentsEnd = append(segmentsEnd, strings.Join(rows[:], "\n")+"\n")
	}
	return strings.Join(segmentsEnd, "")
}
