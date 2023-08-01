package main

import "fmt"

func main() {
	type NoKTP string // contoh membuat type data NoKTP sebagai String
	type Married bool

	var noKTPEko NoKTP = "1231232132132"
	var marriedStatus Married = true

	fmt.Println(noKTPEko)
	fmt.Println(marriedStatus)
}
