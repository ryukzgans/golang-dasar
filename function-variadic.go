package main

import "fmt"

// varadic itu kyk args di python, dapat menerima banyak param
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}

	return total
}

func main() {
	// set param secara langsung
	numbers := sumAll(10, 20, 30, 40, 50)
	fmt.Println(numbers) // 150

	// param melalui slice
	sliceNumbers := []int{10, 20, 30, 40, 50}
	totalSlice := sumAll(sliceNumbers...)
	fmt.Println(totalSlice) // 150
}
