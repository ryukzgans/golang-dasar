package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// anonymous function with variable
	blackList1 := func(name string) bool {
		return name == "admin" // jika nama == anjing, maka kembaliannya adalah true
	}
	registerUser("Ilham", blackList1)
	registerUser("admin", blackList1)

	// anonymous function lgsung
	registerUser("administrator", func(name string) bool {
		return name == "administrator"
	})
}
