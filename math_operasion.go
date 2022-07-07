package main

import "fmt"

func main() {
	var total = 10 + 10
	fmt.Println(total)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	//Augmented Assignment:
	a += 100
	b -= 5
	fmt.Println(a)
	fmt.Println(b)

	//Unary Operator
	//contoh : ++ -- !
	a++
	b--
	fmt.Println(a)
	fmt.Println(b)

}
