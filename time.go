package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Full:", now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())

	date := time.Date(2023, 07, 30, 19, 40, 10, 10, time.Local)
	fmt.Println(date)

	layout := "02-Jan-2006"
	parse, _ := time.Parse(layout, "23-Jul-2023")
	fmt.Println(parse)
}
