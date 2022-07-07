package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var nomorKtpMiraj NoKTP = "32203030303"
	fmt.Println(nomorKtpMiraj)

	var marriedStatus Married = true
	fmt.Println(marriedStatus)
}
