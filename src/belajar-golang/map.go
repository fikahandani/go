package main

import "fmt"

func main() {
	fmt.Println("Belajar penggunaan fungsi Map")
	var name map[string]int
	name = map[string]int{}
	name["Fika"] = 100
	name["Benny"] = 90
	name["Kevin"] = 95
	fmt.Println("Fika", name["Fika"])
	fmt.Println("Benny", name["Benny"])
	fmt.Println("Andre", name["Andre"])
	//nilai map terakhir akan bernilai 0 karena tidak ditemukan pada data yang diinisialisasi

	//map dengan menggunakan perulangan
	var month = map[string]int{
		"Januari":   1,
		"Februari":  2,
		"Maret":     3,
		"April":     4,
		"Mei":       5,
		"Juni":      6,
		"Juli":      7,
		"Agustus":   8,
		"September": 9,
		"Oktober":   10,
		"November":  11,
		"Desember":  12,
	}
	for key, val := range month {
		fmt.Println(key, "   \t", val)
	}
	//nilai kembalian yang dihasilkan tidak selalu urut alias random

	//cara menghapus nilai map
	var mobil = map[string]int{
		"Avanza":  100,
		"Xenia":   98,
		"Mazda":   100,
		"Jazz":    97,
		"Mobilio": 89,
	}
	fmt.Println(len(mobil))
	fmt.Println(mobil)
	delete(mobil, "Mazda")
	fmt.Println(mobil)
}
