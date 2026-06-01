package main

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	segments := SplitInput(input)
	result := ""
	for i, segment := range segments {
		if segment == "" {
			if i < len(segments)-1 {
				result += "\n"
			}
		} else {
			rows := RenderLine(segment, banner)
			for _, row := range rows {
				result += row + "\n"
			}
		}
	}
	return result
}
