package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		https://pkg.go.dev/strings
	*/

	fmt.Println(strings.Trim(" Ilham God ", " "))                               // menghapus awal dan akhir cutset
	fmt.Println(strings.Split("Ilham Kurniawan Tarmiji", " "))                  // split spasi pada kalimat
	fmt.Println(strings.Count("aku adalah raja", "a"))                          // menghitung jumlah a pada kalimat
	fmt.Println(strings.ToUpper("ini semua huruf besar"))                       // mengubah seluruh kata menjadi uppercase
	fmt.Println(strings.ToLower("INI SEMUA HURUF KECIL"))                       // mengubah seluruh kata menjadi lowercase
	fmt.Println(strings.Contains("Ilham God adalah seorang dewa", "dewa"))      // mengecek apakah "dewa" ada di kalimat tersebut
	fmt.Println(strings.ReplaceAll("ilham ilham ilham ilham", "ilham", "budi")) // mereplace semua kata ilham menjadi budi
}
