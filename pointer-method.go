package main

import "fmt"

type Man struct {
	Name string
}

// seperti sebelumnya memakai pointer cukup dengan menambahkan * pada awal typeDatanya
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	budi := Man{Name: "Budi Tarmiji"}
	budi.Married()

	fmt.Println(budi.Name)
}
