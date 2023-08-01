package main

import "fmt"

type Address struct {
	City, Province, Country string
}

/*
	jadi penggunaan pointer pada function itu seperti ini
	jika anda ingin mengubah data asli pada variable alamat dengan typedata address
	maka dibutuhkan pointer, jika tidak data tersebut tidak akan berubah
	func ChangeAddressCountry(address *Address) {}
*/

func ChangeAddressCountry(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	alamat := Address{
		City:     "Medan",
		Province: "Sumatera Utara",
	}

	// dan juga disini wajib diawali dengan &
	ChangeAddressCountry(&alamat)
	fmt.Println(alamat) // {Medan Sumatera Utara Indonesia}
}
