package main

import "fmt"

func main() {
	//deklarasi array basic
	var name [4]string
	name[0] = "Fika"
	name[1] = "Benny"
	name[2] = "Andre"
	name[3] = "Kevin"

	fmt.Println(name[0], name[1], name[2], name[3])

	//pengisian dan menghitung jumlah array
	var fruits = [4]string{"apple", "orange", "banana", "grape"}
	//menghitung jumlah array dengan menggunakan fungsi len
	fmt.Println("Jumlah ", len(fruits))
	//menampilkan data apa saja yang ada pada array
	fmt.Println("Data buat", fruits)

	//inisialisasi nilai awal array tanpa menentukan jumlah index array
	var number = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Jumlah angka", len(number))
	fmt.Println("Data angka", number)

	//array multidimensi
	var numbers1 = [2][3]int{[3]int{1, 2, 3}, [3]int{3, 2, 1}}
	var numbers2 = [2][3]int{[3]int{4, 5, 6}, [3]int{6, 5, 4}}
	fmt.Println("Array pertama", numbers1)
	fmt.Println("Array kedua", numbers2)

	//alokasi array menggunakan keyword make
	var animals = make([]string, 2)
	animals[0] = "Lion"
	animals[1] = "Monkey"
	fmt.Println(animals)
}
