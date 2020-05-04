package main

import "fmt"

func main() {
	fmt.Println("Deteksi Keberadaan Item dengan Key Tertentu")

	var month = map[string]int{"Januari": 90, "Februari": 100, "Maret": 99}
	var value, isExist = month["Mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Data doesn't exist")
	}
}
