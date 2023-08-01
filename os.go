package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		https://pkg.go.dev/os
		untuk lebih lengkapnya
	*/

	args := os.Args // go run os.go argument1 argument2 argument3
	for i, val := range args {
		fmt.Println("Argument", i, "adalah", val)
	}

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", name)
	} else {
		fmt.Println("Error", err)
	}

}
