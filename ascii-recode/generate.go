package main

func GeneratePattern(m rune) []string {
	ways := map[rune][]string {
		'A': {
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
		},

		'B': {
			" #### ",
			"    # ",
			"   #  ",
			"  #   ",
			" #    ",
			" #    ",
			" #### ",
			"      ",
		},

	}
	if m < 'A' || m > 'Z' {
			return []string{}
		}

	end, good := ways[m]
	if !good {
		return []string{
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
		}
	}
	return end
}