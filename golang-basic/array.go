package main

import "fmt"

func main() {
	// Array dengan ukuran tetap [3]string
	var names = [3]string{
		"john",
		"doe",
		"wick",
	}

	fmt.Println(names)
	// Mengakses elemen array dengan index
	fmt.Println(names[0], names[1], names[2])
	// Menghitung panjang array
	fmt.Println(len(names))

	// Array dengan ukuran otomatis [...]
	var values = [...]int{
		90,
		80,
		95,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}
