package main

import "fmt"

func main() {
	fmt.Println("Belajar menggunakan fungsi copy untuk mengcopy array pada slice")

	dst := make([]string, 3)
	src := []string{"watermelon", "apple", "orange", "papaya"}
	n := copy(dst, src)
	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
	//fungsi copy mengembalikan nilai angka berdasarkan hasil dari jumlah elemen yang tercopy
}
