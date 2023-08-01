package main

import "fmt"

func main() {
	// slice array = pointer,length,capacity

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]

	// function yang ada di slice
	// len(slice), cap(slice), append(slice, data)

	fmt.Println(slice1)
	fmt.Println(len(slice1)) // length
	fmt.Println(cap(slice1)) // capacity

	// make a new slice
	// make([]typeData, length, capacity)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ilham"
	newSlice[1] = "Kurniawan"

	fmt.Println("---------------")
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	// copy(destination, source)
	fmt.Println("---------------")
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// array vs slice
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
