package main

import "fmt"

/*
Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
Sebuah interface berisikan definisi-definisi method
Biasanya interface digunakan sebagai kontrak
*/

type HasName interface {
	GetName() string
	/*
		jadi kontrakny itu berisi
		method GetName() dengan output harus string
	*/
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// kemudian agar interface bisa digunakan kita butuh struct sebagai kontraknya
type Person struct {
	Name string
}

// membuat function sesuai dgn interface yaitu GetName() dengan output berupa string
func (person Person) GetName() string {
	return person.Name
}

// contoh membuat struct yang berbeda tapi ttp bisa dipanggil menggunakan interface
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	// deklarasi
	ilham := Person{Name: "Ilham"}
	sayHello(ilham)

	kucing := Animal{Name: "BUDI"}
	sayHello(kucing)
}
