package main

import "fmt"

func main() {
	fmt.Println("Belajar Constant, nilai yang tidak bisa ubah")

	const firstname = "Muhammad"
	const lastname string = "Muhammad"
	const value = 100

	//muliple constant:
	const (
		alamat = "selaras"
		city   = "depok"
	)

	fmt.Println(alamat, city)

}
