package main

import "fmt"

type Address struct{
	City, Province, Country string
}

func main(){
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address1

	// Menggunakan asterisk operator (*) untuk mengakses nilai yang ditunjuk pointer
	// *address2 berarti "nilai yang ditunjuk oleh address2"
	// Operasi ini mengganti seluruh nilai address1 dengan struct Address baru
	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"}

	// address1 sekarang berisi nilai yang baru karena diubah melalui pointer
	fmt.Println(address1) // Output: {Bandung Jawa Barat Indonesia}
	
	// address2 tetap menunjuk ke alamat memori yang sama, tapi nilai sudah berubah
	fmt.Println(address2) // Output: &{Bandung Jawa Barat Indonesia}
}