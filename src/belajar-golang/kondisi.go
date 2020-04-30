package main

import "fmt"

func main() {
	var point = 8840.0
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var a = 5.0
	if a == 5 {
		fmt.Println("This is A")
	} else {
		fmt.Println("This is not A")
	}

	//perulangan
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	//perulangan tanpa argumen
	var i = 0
	for {
		fmt.Println("Angka Tanpa Argumen", i)
		i++
		if i == 5 {
			break
		}
	}
}
