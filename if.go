package main

import "fmt"

func main() {
	var name = "Ilham"

	if name == "Ilham" {
		fmt.Println("Hello Ilham")
	} else if name == "Kurniawan" {
		fmt.Println("Hello Kurniawan")
	} else {
		fmt.Println("Hello Sir")
	}

	// if short statement
	// contoh sebelum
	// lengthName := len(name)
	if lengthName := len(name); lengthName > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
