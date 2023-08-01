package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	// assign variable sebagai function
	sayGoodBye := getGoodBye

	// tinggal panggil variable beserta param
	fmt.Println(sayGoodBye("Ilham")) // Good Bye Ilham
}
