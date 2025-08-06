package main

import "fmt"

func main() {
	// Menghitung panjang string
	fmt.Println(len("Hello, World!"))
	
	// Mengakses karakter pada index tertentu (byte value)
	fmt.Println("Hello, World!"[1])
	
	// String slicing - mengambil substring
	fmt.Println("Hello, World!"[0:5])
}