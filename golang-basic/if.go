package main

import "fmt"

func main() {
	name := "John Doe"

	// Contoh if-else statement dasar
	// Mengecek kondisi dengan operator perbandingan ==
	if name == "John Doe" {
		fmt.Println("Hello, John Doe")
	} else if name == "Jane Doe" {
		fmt.Println("Hello, Jane Doe")
	} else {
		fmt.Println("Hi, there")
	}

	// If statement dengan short variable declaration
	// Variable 'length' hanya tersedia dalam scope if-else ini
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}