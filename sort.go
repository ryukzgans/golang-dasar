package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

// inisialisasi UserSlice
type UserSlice []User

// membuat kontrak Len() terhadap UserSlice
func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

// membuat kontrak Less() terhadap UserSlice
func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

// membuat kontrak Less() terhadap UserSlice
func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	dataUser := []User{
		{"Budi", 20},
		{"Ilham", 19},
		{"Kurniawan", 21},
		{"Tarmiji", 18},
	}

	fmt.Println(dataUser) // sebelum di sort
	sort.Sort(UserSlice(dataUser))
	fmt.Println(dataUser) // sesudah di sort
}
