package main

import "fmt"

func main() {
	// for init; kondisi; post {}
	for number := 1; number <= 10; number++ {
		fmt.Println("Perulangan ke", number)
	}

	// for range
	slice1 := []string{"ilham", "kurniawan", "ahmad", "biji", "ayam"}
	map1 := map[string]string{"firstName": "Ilham", "lastName": "Kurniawan"}

	// contoh pada slice / array
	for i, val := range slice1 {
		fmt.Println("Index", i, "=", val)
	}

	// contoh pada map
	for k, v := range map1 {
		fmt.Println(k, "=", v)
	}

	// jika tidak ingin index tersebut digunakan bisa mengganti symbolnya menjadi _
	// for _, val := range slice1 {}
}
