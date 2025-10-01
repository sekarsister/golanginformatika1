package main

import "fmt"

func main() {
	var r int
	fmt.Print("Masukkan jari-jari (r): ")
	fmt.Scan(&r)

	var luas float64
	if r%7 == 0 {
		luas = float64(22*r*r) / 7.0
	} else {
		luas = 3.14 * float64(r*r)
	}

	fmt.Printf("Luas lingkaran = %.2f\n", luas)
}
