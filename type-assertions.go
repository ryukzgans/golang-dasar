package main

import "fmt"

/*
Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
*/

func random() interface{} {
	return "OK"
}

func main() {
	random := random()

	// harus dipastikan sebelum mengganti tipe data, kembalian dari random() harus merupakan string
	// jika tidak maka akan error panic
	randomString := random.(string)
	fmt.Println(randomString)

	// contoh error panic
	// randomInt := random.(int)
	// fmt.Println(randomInt)

	// contoh lebih praktis bisa menggunakan switch
	switch value := random.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	case bool:
		fmt.Println("Value", value, "is bool")
	default:
		fmt.Println("Unknown")
	}

}
