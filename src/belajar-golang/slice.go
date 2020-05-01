package main

import "fmt"

func main() {
	//deklarasi slice
	var fruits = []string{"apple", "banana", "grape", "mango"}
	fmt.Println(fruits[0])
	//perbedaan dengan menggunakan array adalah pada deklarasi jumlah elemen
	//jika tidak ada deklarasi jumlah elemen maka variable tersebut adalah slice

	//hubungan array dengan slice
	var animals = []string{"lion", "tiger", "ant", "monkey"}
	var newAnimal = animals[0:2] //artinya mengakses data array dari index ke 0 sampai sebelum index ke 2
	fmt.Println(newAnimal)
}
