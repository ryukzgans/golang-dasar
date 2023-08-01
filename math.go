package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(2.4)) // 2 || Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat
	fmt.Println(math.Round(2.6)) // 3 || Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat

	fmt.Println(math.Floor(2.7)) // 2 || Membulatkan float64 kebawah
	fmt.Println(math.Ceil(2.2))  // 3 || Membulatkan float64 keatas

	fmt.Println(math.Max(5.3, 6.4)) // 6.4 || Mengembalikan nilai float64 paling besar
	fmt.Println(math.Min(5.3, 6.4)) // 5.3 || Mengembalikan nilai float64 paling kecil
}
