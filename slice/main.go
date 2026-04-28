package main

import (
	"fmt"
	"strings"
)

func main() {
	slice()
	nums := []int{10, 20, 30}

	for i, v := range nums {
		fmt.Println(i, v)
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	words := []string{"Hello", "World"}
	fmt.Println(words[0]) // Hello

	//Reading a file and turning it into lines:\
	content := "line1\nline2\nline3"

	lines := strings.Split(content, "\n")

	fmt.Println(lines)
	// ["line1", "line2", "line3"]

}
