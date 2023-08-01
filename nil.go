package main

import "fmt"

// nil itu sama sprti null = data kosong
// Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel

// contoh membuat function newMap
func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{"name": name}
	}
}

func main() {
	newMap := newMap("")
	fmt.Println(newMap)

	// contoh singkat penggunaan nil
	if newMap == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(newMap)
	}
}
