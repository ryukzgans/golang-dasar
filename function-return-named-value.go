package main

import "fmt"

// return named value / return named multiple value
func getFullName() (firstName, middleName, lastName string) {
	// assign lgsung
	firstName = "Ilham"
	middleName = "Kurniawan"
	lastName = "Sitompul"

	return // bisa lgsung return gini atau
	// return firstName, middleName, lastName // atau bisa gini
}

func main() {
	fn, mn, ln := getFullName() // nama variable bebas

	fmt.Println(fn, mn, ln)
}
