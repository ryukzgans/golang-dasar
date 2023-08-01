package main

import (
	"flag"
	"fmt"
)

func main() {
	// Package flag berisikan fungsionalitas untuk memparsing command line argument
	host := flag.String("host", "localhost", "put your database host") // flag.String(name, default value, description)
	username := flag.String("user", "root", "put your database user")
	password := flag.String("password", "root@123", "put your database password")

	flag.Parse() // wajib mendeklarasikan flag.Parse() agar argument yg di inisialisasi berjalan

	fmt.Println("Host:", *host)
	fmt.Println("User:", *username)
	fmt.Println("Password:", *password)

	// go run flag.go -host=localhost -user=root -password=root@123
}
