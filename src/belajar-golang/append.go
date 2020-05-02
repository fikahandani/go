package main

import "fmt"

func main() {
	fmt.Println("Belajar menggunakan fungsi Append pada Golang")

	var fruits = []string{"apple", "banana", "mango", "orange"}
	var addFruit = append(fruits, "papaya")
	fmt.Println(fruits)
	fmt.Println(addFruit)

	var aFruit = fruits[0:3]
	var bFruit = append(aFruit, "Dragon Fruit")
	fmt.Println(aFruit)
	fmt.Println(bFruit)
	fmt.Println(fruits)
	//jika nilai len kurang dari nilai cap, maka nilai append akan dianggap sebagai nilai reference
}
