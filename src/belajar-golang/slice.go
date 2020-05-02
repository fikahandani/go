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

	//slice sebagai reference
	var names = []string{"Fika", "Anton", "Dewi", "Dinda"}
	var aName = names[0:3]
	var bName = names[1:4]
	var aaName = aName[1:2]
	var bbName = bName[0:1]

	fmt.Println(names)
	fmt.Println(aName)
	fmt.Println(bName)
	fmt.Println(aaName)
	fmt.Println(bbName)

	//penggunaan fungsi len
	//fungsi len berfungsi untuk menghitung jumlah data slice
	var city = []string{"Malang", "Jakarta", "Surabaya", "Denpasar", "Blitar"}
	fmt.Println(len(city))

	//penggunaan fungsi cap
	//fungsi cap digunakan untuk menghitung kapasitas maksimal sebuah slice
	var country = []string{"Indonesia", "US", "Autralia", "India"}
	var aCountry = country[0:3]
	fmt.Println(len(aCountry))
	fmt.Println(cap(aCountry))
}
