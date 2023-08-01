package main

import "fmt"

// factorial dengan loop
func factorialLoop(value int) int {
	total := 1
	for i := value; i > 0; i-- {
		total *= i
	}
	return total
}

// factorial dengan recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop := factorialLoop(10)
	fmt.Println(loop)

	recursive := factorialRecursive(10)
	fmt.Println(recursive)
}
