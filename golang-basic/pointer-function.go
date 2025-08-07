package main

import "fmt"

type Address struct{
	City, Province, Country string
}

// Function yang menerima parameter pointer ke Address
// Dengan menggunakan pointer, function dapat memodifikasi nilai asli
// Parameter *Address berarti "pointer ke Address"
func ChangeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main(){
	address := Address{}
	
	// Memanggil function dengan mengirim alamat memori address menggunakan &
	// &address berarti "alamat memori dari address"
	ChangeCountryToIndonesia(&address)

	// address sekarang memiliki Country = "Indonesia" karena diubah melalui pointer
	fmt.Println(address) // Output: { Indonesia}
}