package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func main() {
	for i := 1; i <= 5; i++ {
		sayHello()
	}
}
