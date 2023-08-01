package main

import (
	"container/list"
	"fmt"
)

// Package container/list adalah implementasi struktur data double linked list di Go-Lang

func main() {
	data := list.New()

	data.PushBack("Ilham")
	data.PushBack("Kurniawan")
	data.PushBack("Tarmiji")
	data.PushBack("God")

	fmt.Println(data.Front().Value)        // ngambil yg pertama
	fmt.Println(data.Back().Value)         // ngambil yg terakhir
	fmt.Println(data.Front().Next().Value) // ngambil data setelah data pertama
	fmt.Println(data.Back().Prev().Value)  // ngambil data sebelum data terakhir

	// contoh loop untuk ngakses data
	// looping dari data awal ke akhir
	fmt.Println("\nData awal ke akhir:")
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// looping dari data akhir ke awal
	fmt.Println("\nData akhir ke awal:")
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
