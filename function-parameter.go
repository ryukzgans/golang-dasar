package main

import "fmt"

func sayHelloTo(firstName string, lastName string, age int) {
	fmt.Println("FirstName:", firstName)
	fmt.Println("LastName:", lastName)
	fmt.Println("Age:", age)
}

func testFunc(firstName, lastName string) {
	fmt.Println(firstName, lastName)
}

func main() {
	sayHelloTo("Ilham", "Kurniawan", 19)
}
