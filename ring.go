package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
 Package container/ring adalah implementasi struktur data circular list
 Circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)
*/

func main() {
	data := ring.New(5)

	// data.Value = "Ilham" // cara memasuki datanya

	// // tetapi kita bagaimana dengan memasuki data selanjutnya ? kita harus membuat variable baru lagi
	// data2 := data.Next()
	// data2.Value = "Kurniawan"

	// cara simple dalam memasuki value ke dalam data ring, kita dpat menggunakan iterasi looping
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// fmt.Println(*data) // tidak dapat lgsung dipanggil dgn println
	// berikut cara memanggilnya secara iterasi
	data.Do(func(val interface{}) {
		fmt.Println(val)
	})
	/*
		Value 0
		Value 1
		Value 2
		Value 3
		Value 4
	*/
}
