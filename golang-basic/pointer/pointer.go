package main

import "fmt"

type Address struct{
	City, Province, Country string
}

func main(){
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	
	// address2 adalah pointer yang menunjuk ke alamat memori address1
	// Menggunakan operator & untuk mendapatkan alamat memori
	address2 := &address1

	// Mengubah nilai City melalui pointer address2
	// Perubahan ini akan mempengaruhi address1 karena keduanya menunjuk ke memori yang sama
	address2.City = "Bandung"

	fmt.Println(address1) // Output: {Bandung DKI Jakarta Indonesia}
	
	// address2 menampilkan alamat memori, bukan nilai langsung
	fmt.Println(address2) // Output: &{Bandung DKI Jakarta Indonesia}
}