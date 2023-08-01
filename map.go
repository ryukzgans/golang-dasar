package main

import "fmt"

func main() {
	// var name = map[typeData key]typeData value{}

	person := map[string]string{
		"name":    "Ilham Kurniawan",
		"address": "Jalan Imam Bonjol",
	}

	// menambah key dan value baru
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// function map
	// len(map), delete(map, key), make(map[typeData]typeData)

	// setup map
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Ilham"
	book["ups"] = "Salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(len(book))
	fmt.Println(book)

}
