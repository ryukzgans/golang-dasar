package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// parameter customer dengan type struct
func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func (cus Customer) sayWelcome(name string) {
	fmt.Println("Welcome", name, "Halo nama saya", cus.Name)
}

func main() {
	// jadi budi itu memiliki sebuah function sayHello() dan sayWelcome() goks sih mirip class gt
	budi := Customer{Name: "Budi"}
	budi.sayHello()
	budi.sayWelcome("Ilham")
}
