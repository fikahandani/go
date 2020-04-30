package main

import "fmt"

func main() {
	//pointer
	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)
	fmt.Println("Fika Handani")
}
