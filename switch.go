package main

import "fmt"

func main() {
	name := "Janc"

	switch name {
	case "Kurniawan":
		fmt.Println("Hello Kurniawan")
	case "Ilham":
		fmt.Println("Hello Ilham")
	default:
		fmt.Println("Hello World")
	}

	// switch short
	switch lengthName := len(name); lengthName > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	lengthName := len(name)
	switch {
	case lengthName > 5:
		fmt.Println("Nama terlalu panjang")
	case lengthName < 5:
		fmt.Println("Nama sudah benar")
	}
}
