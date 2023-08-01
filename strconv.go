package main

import (
	"fmt"
	"strconv"
)

/*
strconv.parseBool(string) => Mengubah string ke bool
strconv.parseFloat(string) => Mengubah string ke float
strconv.parseInt(string) => Mengubah string ke int64
strconv.FormatBool(bool) => Mengubah bool ke string
strconv.FormatFloat(float, … ) => Mengubah float64 ke string
strconv.FormatInt(int, … ) => Mengubah int64 ke string
*/

func main() {
	// https://pkg.go.dev/strconv

	boolean, err := strconv.ParseBool("false") // konversi string ke bool
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	/*
		base number > binary=2, octal=8, desimal=10, hexadesimal=16
		bitsize > 32, 64
	*/
	num, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(num)
	} else {
		fmt.Println(err.Error())
	}

	// contoh lebih simple konversi string to int atau sebaliknya, dgn menggunakan Atoi / Itoa
	val1, _ := strconv.Atoi("2000000") // string to integer
	fmt.Println(val1)

	val2 := strconv.Itoa(2000000) // integer to string
	fmt.Println(val2)
}
