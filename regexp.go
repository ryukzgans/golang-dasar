package main

import (
	"fmt"
	"regexp"
)

// Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
/*
	regexp.MustCompile(string) => Membuat Regexp
	Regexp.MatchString(string) bool => Mengecek apakah Regexp match dengan string
	Regexp.FindAllString(string, max) => Mencari string yang match dengan maximum jumlah hasil
*/

func main() {
	regex := regexp.MustCompile("e([a-z])o") // e([a-z])o itu berarti harus diawali i dan diakhiri m, kemudian di tengah a-z

	fmt.Println(regex.MatchString("eko")) // true
	fmt.Println(regex.MatchString("eto")) // true
	fmt.Println(regex.MatchString("eKo")) // false

	value := regex.FindAllString("eko edo egi ego e1o eto", -1)
	fmt.Println(value)
}
