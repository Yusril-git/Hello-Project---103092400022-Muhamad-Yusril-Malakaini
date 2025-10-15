package main

import "fmt"

func main() {
	var jarak,kecepatan int
	var waktu, posisi int

	fmt.Println("masukan jarak, kecepatan dan waktu")
	fmt.Scanln(&jarak, &kecepatan, &waktu)

	posisi = (kecepatan * waktu) + jarak
	fmt.Println(posisi)
}