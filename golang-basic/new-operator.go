package main

import "fmt"

type Address struct{
	City, Province, Country string
}

func main(){
	// Menggunakan new() untuk membuat pointer ke struct Address baru
	// new(Address) mengalokasikan memori dan mengembalikan pointer ke zero value
	// Zero value untuk string adalah "" (string kosong)
	var address1 *Address = new(Address)
	
	// address2 adalah pointer yang menunjuk ke alamat memori yang sama dengan address1
	// Kedua pointer ini berbagi objek yang sama di memori
	var address2 *Address = address1

	address2.Country = "Indonesia"

	fmt.Println(address1) // Output: &{ Indonesia}
	fmt.Println(address2) // Output: &{ Indonesia}
}