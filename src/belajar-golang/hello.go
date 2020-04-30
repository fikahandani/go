package main

import "fmt"

func main() {

	//output print text standard
	fmt.Println("Kenalan sama Golang")
	fmt.Println("Asik juga ternyata")

	//deklarasi variable string
	var firstName string = "Fika"
	var lastName string = "Handani"
	fmt.Printf("Hello %s %s\n", firstName, lastName)
	fmt.Println("Hello ", firstName, lastName+" Sayang kamu")

	//deklarasi text tanpa var
	var namaDepan string = "Dia"
	namaBelakang := "Sebenarnya"
	fmt.Println("Hello", namaDepan, namaBelakang+" Sayang Fika")

	//multiple declaration
	var depan string = "Seseorang"
	belakang := "Annoying"
	status := "ngejar-ngejar"
	fmt.Println("Hello", depan, belakang, status)

	//multi variable
	var x, y, z string
	x, y, z = "Fika", "Sayang", "Kamu"
	a, b, c := "Kamu", "Juga", "Sayang Fika"
	fmt.Println("Harus begini ya", x, y, z, a, b, c)
}
