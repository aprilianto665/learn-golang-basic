package main

import "fmt"

func main() {
	// Closure adalah fungsi yang dapat mengakses variabel dari scope luar
	// Variabel counter dideklarasikan di scope main function
	counter := 0

	// increment adalah anonymous function (closure) yang "menangkap" variabel counter
	// Meskipun counter berada di luar scope function increment, tetap bisa diakses
	increment := func() {
		fmt.Println("Increment")
		counter++ // Mengubah nilai counter dari scope luar
	}

	// Memanggil closure increment dua kali
	increment() // counter menjadi 1
	increment() // counter menjadi 2

	// Menampilkan nilai akhir counter yang telah dimodifikasi oleh closure
	fmt.Println(counter) // Output: 2
}