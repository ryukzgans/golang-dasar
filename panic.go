package main

import "fmt"

// Panic function adalah function yang bisa kita gunakan untuk menghentikan program
// Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
// Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

func runApp(error bool) {
	defer endApp() // akan selalu berjalan
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi telah berjalan")
}

func endApp() {
	fmt.Println("Aplikasi telah selesai")
}

func main() {
	runApp(true)
	// Aplikasi telah selesai > defer tetap berjalan
	// panic: APLIKASI ERROR > panic, terjadi error
}
