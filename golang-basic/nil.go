package main

import "fmt"

// KONSEP NIL DI GO:
// nil adalah zero value untuk reference types: pointer, slice, map, channel, function, interface
// nil berbeda dengan empty value - nil berarti "tidak ada referensi ke memori"
// Mengakses nil map/slice akan menyebabkan panic, jadi selalu lakukan nil check

// Function yang mendemonstrasikan penggunaan nil di Go
func NewMap(name string) map[string]string {
	// Mengecek apakah parameter name adalah string kosong
	if name == "" {
		return nil // Mengembalikan nil jika name kosong
	} else {
		// Membuat dan mengembalikan map baru dengan key "name"
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// Memanggil NewMap dengan string kosong, akan mengembalikan nil
	data := NewMap("")
	
	// Mengecek apakah data bernilai nil
	// nil check adalah praktik umum di Go untuk menghindari panic
	if data == nil {
		fmt.Println("Data Kosong") // Output: Data Kosong
	} else {
		fmt.Println(data) // Akan mencetak map jika tidak nil
	}
}