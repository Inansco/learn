package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("use: go run . input.txt output.txt")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error")
		return
	}

	data := Processor(string(file))

	err = os.WriteFile(output, []byte(data), 0644)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println("It is done!!!!!!!")
}
