package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 20, 30}

	for _, v := range nums {
		fmt.Println(v)
	}
	for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}
}
