package main

import "fmt"

// func name(name typeData) return typeData {}
func sayHello(name string) string {
	if name != "" {
		return "Hello " + name
	} else {
		return "Hello Bro"
	}
}

func main() {
	res := sayHello("")
	fmt.Println(res)
	fmt.Println(sayHello("Kurniawan"))

}
