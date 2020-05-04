package main

import "fmt"

func main() {
	fmt.Println("Kombinasi Slice dan Map")

	//untuk isi data yang sama
	var mahasiswa = []map[string]string{
		{"Nama": "Fika", "Nim": "311310009"},
		{"Nama": "Benny", "Nim": "311310004"},
		{"Nama": "Kevin", "Nim": "311310012"},
	}
	for _, mhs := range mahasiswa {
		fmt.Println(mhs["Nama"], mhs["Nim"])
	}

	//kombinasi slice dan map untuk isi data yang berbeda
	var dosen = []map[string]string{
		{"Nama": "Windra Swastika", "Gender": "Male", "Status": "Dosen"},
		{"Address": "Malang", "Nomor": "120"},
		{"Kantor": "Universitas Ma Chung"},
		{"Nama": "Ketrilia", "Gender": "Female", "Status": "Dosen"},
		{"Address": "Surabaya", "Nomor": "130"},
		{"Kantor": "Universitas Ma Chung"},
	}
	for _, dsn := range dosen {
		fmt.Println(dsn["Nama"])
	}
}
