package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Funsi Variadic digunakan untuk memasukan nilai argumen pada sebuah parameter yang tidak terbatas jumlahnya")
	var avg = calculate(2, 3, 2, 5, 6, 10, 11, 23, 21, 6)
	var msg = fmt.Sprintf("Average : %.2f", avg)
	fmt.Println(msg)

	fmt.Println("Jika menggunakan slice sebagai nilai argumen pada parameter")
	var number = []int{2, 3, 3, 5, 6, 1, 4, 9, 2, 1}
	var avg2 = calculate(number...)
	var msg2 = fmt.Sprintf("Avegare 2 : %.2f", avg2)
	fmt.Println(msg2)

	fmt.Println("Menggabungkan parameter biasa dengan variadic")
	fmt.Println("parameter variadic harus berada pada posisi akhir")
	data("Fika", "Makan", "Tidur", "Ngoding ga jelas")

}
func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers)) //casting variable dari int menjadi float
	return avg
}
func data(name string, hobbies ...string) {
	var data = strings.Join(hobbies, ", ")
	fmt.Printf("Hello, My Name is %s\n", name)
	fmt.Printf("And My Hobbies is %s\n", data)
}
