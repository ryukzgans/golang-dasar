package main

import "fmt"

// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

// intinya recover digunakan untuk menangkap data panic, jika kita meletakannya pada defer

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println(message)
		// jika message tidak nil/null
		// keluar program akan jadi "APLIKASI ERROR" > sama kyk panic message
		// tetapi program tetap berjalan
	}
	fmt.Println("Aplikasi Selesai")
}

func main() {
	runApp(true)
	fmt.Println("Ilham")
}
