package main

import "fmt"

func main() {
	var name string // normal variable

	name = "Ilham Kurniawan"
	fmt.Println(name)

	name = "Ilham God"
	fmt.Println(name)

	var age = 30 // tanpa typedata
	fmt.Println(age)

	nama := "Ilham Gunawan" // tanpa typedata dan assign var
	fmt.Println(nama)

	// multiple variable
	var (
		firstName = "Ilham"
		lastName  = "Kurniawan"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
