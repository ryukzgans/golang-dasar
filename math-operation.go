package main

import "fmt"

func main() {
	result := 10 + 10
	fmt.Println(result)

	a := 10
	b := 10
	c := 10
	fmt.Println(a + b + c)

	// augmented operator
	result += 10 // i = i + 10
	fmt.Println(result)

	// unary operator
	result++ // i = i + 1
	fmt.Println(result)

	negative := -100
	positive := +100
	fmt.Println(negative)
	fmt.Println(positive)
}
