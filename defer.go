package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	// Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
	// Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApplication(0) // got an error panic runtime: divide by zero
	// Run Application
	// Selesai memanggil function
	// setelahnya baru error
}
