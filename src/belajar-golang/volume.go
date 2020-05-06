package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Funsi Random untuk menghitung volume bangun ruang")
	rand.Seed(time.Now().Unix())
	var kubus = kubus(12)
	fmt.Println("Volume Kubus ", kubus)
	var tabung = tabung(10.0, 10.5)
	fmt.Println("Volume Tabung ", tabung)
	var balok = balok(10, 8, 12)
	fmt.Println("Volume Balok ", balok)
}

func kubus(sisi int) int {
	var volumeKubus = sisi * sisi * sisi
	return volumeKubus
}

func tabung(t, r float32) float32 {
	var volumeTabung float32 = 3.14 * r * r * t
	return volumeTabung
}

func balok(panjang, lebar, tinggi int) int {
	var volumeBalok = panjang * lebar * tinggi
	return volumeBalok
}
