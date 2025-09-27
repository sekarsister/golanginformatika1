package main

import (
	"fmt"
	"math"
)

func main() {
	var pilihan int
	fmt.Println("=== PROGRAM MENGHITUNG LUAS DAN KELILING ===")
	fmt.Println("1. Segitiga")
	fmt.Println("2. Lingkaran")
	fmt.Println("3. Persegi Panjang")
	fmt.Println("4. Persegi")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih bangun datar (1-5): ")
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		segitiga()
	case 2:
		lingkaran()
	case 3:
		persegiPanjang()
	case 4:
		persegi()
	case 5:
		fmt.Println("Terima kasih telah menggunakan program!")
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}

func segitiga() {
	var alas, tinggi, sisi1, sisi2, sisi3 float64
	fmt.Println("\n=== SEGITIGA ===")
	fmt.Print("Masukkan alas segitiga: ")
	fmt.Scanln(&alas)
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scanln(&tinggi)
	fmt.Print("Masukkan sisi 1 segitiga: ")
	fmt.Scanln(&sisi1)
	fmt.Print("Masukkan sisi 2 segitiga: ")
	fmt.Scanln(&sisi2)
	fmt.Print("Masukkan sisi 3 segitiga: ")
	fmt.Scanln(&sisi3)

	luas := 0.5 * alas * tinggi
	keliling := sisi1 + sisi2 + sisi3

	fmt.Printf("\nHasil Perhitungan Segitiga:\n")
	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n", keliling)
}

func lingkaran() {
	var jariJari float64
	fmt.Println("\n=== LINGKARAN ===")
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&jariJari)

	luas := math.Pi * jariJari * jariJari
	keliling := 2 * math.Pi * jariJari

	fmt.Printf("\nHasil Perhitungan Lingkaran:\n")
	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n", keliling)
}

func persegiPanjang() {
	var panjang, lebar float64
	fmt.Println("\n=== PERSEGI PANJANG ===")
	fmt.Print("Masukkan panjang: ")
	fmt.Scanln(&panjang)
	fmt.Print("Masukkan lebar: ")
	fmt.Scanln(&lebar)

	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Printf("\nHasil Perhitungan Persegi Panjang:\n")
	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n", keliling)
}

func persegi() {
	var sisi float64
	fmt.Println("\n=== PERSEGI ===")
	fmt.Print("Masukkan sisi persegi: ")
	fmt.Scanln(&sisi)

	luas := sisi * sisi
	keliling := 4 * sisi

	fmt.Printf("\nHasil Perhitungan Persegi:\n")
	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n", keliling)
}
