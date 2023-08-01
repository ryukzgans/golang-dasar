package main

import (
	"errors"
	"fmt"
)

// misal penggunaan error
func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL tidak diperbolehkan")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, err := Pembagian(10, 10)

	if err == nil {
		fmt.Println("Hasil:", result)
	} else {
		fmt.Println("Error:", err)
	}
}
