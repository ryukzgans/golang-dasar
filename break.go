package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke", i)
		if i == 7 {
			break
		}
	}
}
