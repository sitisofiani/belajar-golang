// konversi tipe data

package main

import "fmt"

func main() {
	var nilai32 int32 = 200000
	var nilai64 int64 = int64(nilai32)

	var name = "Samsul"
	var e byte = name[0]
	var eString string = string(e)
	

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(name)
	fmt.Println(eString)

}