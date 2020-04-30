package main

import "fmt"

func main() {

	//integer value menggunakan println
	var number int = 5
	fmt.Println("This number", number)

	//integer value menggunakan prntf
	var number2 int = 7
	fmt.Printf("This number is %d\n", number2)

	//float number
	var floatNumber float32 = 3.14
	fmt.Println("This float number", floatNumber)

	//float number pakai printf
	var floatNumber2 float64 = 12.98123
	fmt.Printf("This is float 64 number %.2f\n", floatNumber2)

	//boolean value
	var exis bool = true
	fmt.Println("This is boolean value", exis)

	//nilai konstanta
	const name = "Fika Handani"
	fmt.Println("My Name is", name)

	//nilai konstanta integer
	const intNumber = 10
	fmt.Println("Value const int Number", intNumber)
}
