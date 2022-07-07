package main

import "fmt"

func main() {

	var name string

	name = "Muhammad Miraj"
	fmt.Println(name)

	name = "Miraj Muhammad"
	fmt.Println(name)

	//tidak wajib memberikan tipe data jika langsung memberikan value.a
	var alamat = "Graha Selaras"
	fmt.Println(alamat)

	var age int8 = 17
	fmt.Println(age)

	//keyword var tidak wajib tetapi harus langsung memberikan value dengan :=
	country := "indonesia"
	fmt.Println(country)

	//multiple variable:
	var (
		firstname = "muhammad"
		lastname  = "miraj"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
}
