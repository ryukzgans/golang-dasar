package main

import "fmt"

// return multiple value
func getFullName() (string, string) {
	return "Ilham", "Kurniawan"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	firstName1, _ := getFullName() // menggunakan _ untuk menghiraukan return value
	fmt.Println(firstName1)
}
