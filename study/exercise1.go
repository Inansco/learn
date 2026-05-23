package main

import "fmt"

func main() {
	//Exercise 1
	var name = "Monday"
	age := 30
	food := "Beans"

	fmt.Printf("My name is %s \n", name)
	fmt.Printf("I am %d years old \n", age)
	fmt.Printf("My favorite food is %s \n", food)

	//Exercise 2
	var country string = "Nigeria"
	var height float32 = 64.8
	var isStudent bool = true

	fmt.Println(country)
	fmt.Println(height)
	fmt.Println(isStudent)


	//Exercise 3
    x := 10
    y := 5

    fmt.Println(x + y) // = 15
    fmt.Println(x - y) // = 5
    fmt.Println(x * y) // = 50

	//Mini Challenge
	name1 := "John"
	age1 := 22
	country1 := "Nigeria"

	fmt.Println("----- PROFILE -----")
	fmt.Printf("Name: %s\n", name1)
	fmt.Printf("Age: %d\n", age1)
	fmt.Printf("Country: %s\n", country1)
	fmt.Println("-------------------")

}