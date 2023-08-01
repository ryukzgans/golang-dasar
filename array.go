package main

import "fmt"

func main() {
	// manual way
	var names [3]string // create array with 3 items
	names[0] = "Ilham"
	names[1] = "Kurniawan"
	names[2] = "Tarmiji"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// langsung
	var values = [3]int{90, 80, 95}

	fmt.Println(values) // [90 80 95]

	fmt.Println(len(names)) // panjang data
	fmt.Println(len(values))
	fmt.Println()
}
