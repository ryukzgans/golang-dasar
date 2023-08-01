package main

//  Reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan

import (
	"fmt"
	"reflect"
)

type User struct {
	Nickname string `required:"true" max:"10"` // `required:"True"` ini adalah structTag
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("required") == "true" {
			// output dari reflect.ValueOf(data).Field(i).Interface() adalah berupa value terhadap field, contoh ilhamgod
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}

	return true
}

func main() {
	user := User{Nickname: "IlhamGod"}

	userType := reflect.TypeOf(user)

	fmt.Println(userType.NumField())                   // mengambil jumlah field pada struct User: 1 yaitu Nickname
	fmt.Println(userType.Field(0).Name)                // mengambil field pada index 0 (pertama) kemudian mereturn nama fieldnya
	fmt.Println(userType.Field(0).Tag)                 // untuk melihat structTag apa saja yang ada pada data struct tersebut
	fmt.Println(userType.Field(0).Tag.Get("required")) // untuk melihat value structTag pada data struct

	fmt.Println(IsValid(user)) // true

	user2 := User{Nickname: ""}
	fmt.Println(IsValid(user2)) // false (tidak boleh string kosong)
}
