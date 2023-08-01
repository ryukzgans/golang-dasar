package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) string {
	filtering := filter(name)
	return "Hello " + filtering
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

// misal dalam parameter func memiliki paramater yang banyak dll
// maka dapat mengdeklarisikannya ke typedata alias seperti ini
type Filter func(string) string

// maka cukup memanggilanya seperti ini
func sayHelloWithFilter2(name string, filter Filter) string {
	return ""
}

func main() {
	result := sayHelloWithFilter("ilham", spamFilter)
	result1 := sayHelloWithFilter("anjing", spamFilter)
	fmt.Println(result)
	fmt.Println(result1)
}
