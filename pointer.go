package main

import "fmt"

// pointer adalah penggunaan untuk merefer lgsung ke memory dituju
// pointer symbol
/*
	&, *
*/

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Medan", "Sumatera Utara", "Indonesia"}
	// address2 := address1 // misal seperti ini saja, si address2 akan membuat duplikasi dari address1

	// agar address2 merefer lgsung ke address1 maka dibutuhkan pointer
	// penggunaanya adalah dengan = &namavariable
	address2 := &address1 // maka jika diubah data di address2, data pada address1 akan ikut terubah juga

	// contoh penggunaan var
	// var address1 Address = Address{"Medan", "Sumatera Utara", "Indonesia"}
	// var address2 *Address = &address1
	var address3 *Address = &address1

	// misal kita ubah ini maka address1 tidak ikut terubah dikarenakan address2 menduplikasikan bukan merefer
	address2.City = "Bandung"

	// jika seperti ini maka kita akan terlepas dari refer address1 dan akan membuat memory data baru address2
	// jadi address1 tidak akan ikut berubah
	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} // &{Jakarta DKI Jakarta Indonesia}

	// bagaimana jika saya ingin mengubah address1 juga ?
	// yaitu dengan cara memakai operator * pada awalan variable
	// perlu diingat address2 tetap membuat memory baru bukan menggunakan memory yang sama
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// maka secara otomatis address1 akan mengikuti memory pada address2

	address4 := new(Address)
	address4.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
	// address3 juga akan terikut berubah karena address1 sudh merefer ke address2, jadi variable apapun yang merefer ke 1 akan ikut berubah
	fmt.Println(address3)
	fmt.Println(address4)
}
