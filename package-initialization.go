package main

import (
	"belajar-golang/golang-dasar/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result) // "MySQL"
}
