package main

import "fmt"

/*
Interface kosong adalah interface yang tidak memiliki deklarasi method satupun,
hal ini membuat secara otomatis semua tipe data akan menjadi implementasi nya
*/

// jadi return bisa berbentuk typedata apapun
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else if i == 3 {
		return 2.5
	} else if i == 4 {
		return [...]int{1, 2, 3, 4, 5}
	} else {
		return "Ups"
	}

}

func main() {
	ups1 := Ups(1)
	ups2 := Ups(2)
	ups3 := Ups(3)
	ups4 := Ups(4)
	ups5 := Ups(5)

	fmt.Println(ups1)
	fmt.Println(ups2)
	fmt.Println(ups3)
	fmt.Println(ups4)
	fmt.Println(ups5)
}
