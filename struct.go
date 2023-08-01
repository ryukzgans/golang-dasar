package main

import "fmt"

// struct
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	// contoh initial pertama
	var ilham Customer
	ilham.Name = "Ilham Kurniawan"
	ilham.Address = "Jalan Karya Budi"
	ilham.Age = 18
	fmt.Println(ilham)

	// contoh inital kedua
	budi := Customer{
		Name:    "Budi",
		Address: "Jalan Kambing",
		Age:     15,
	}
	fmt.Println(budi)

	// contoh initial ketiga | tidak terlalu di sarankan
	sutomo := Customer{"Sutomo", "Jalan Banteng", 15} // diisi sesuai urutan
	fmt.Println(sutomo)

}
