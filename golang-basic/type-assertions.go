package main

import "fmt"

// TYPE ASSERTIONS DI GO:
// Type assertion digunakan untuk mengekstrak nilai konkret dari interface{} atau any
// Syntax: value.(Type) - akan panic jika type tidak sesuai
// Syntax aman: value, ok := value.(Type) - mengembalikan boolean untuk mengecek keberhasilan
// PERINGATAN: Kode ini akan panic karena mencoba convert string ke int!

func random() any {
	return "OK"
}

func main() {
	result := random()
	
	// TYPE ASSERTION TIDAK AMAN:
	// Menggunakan syntax result.(string) - akan berhasil karena result memang string
	resultString := result.(string)
	fmt.Println(resultString)

	// TYPE ASSERTION BERBAHAYA:
	// Mencoba mengkonversi ke int padahal result adalah string
	// AKAN MENYEBABKAN PANIC! Program akan crash di sini
	resultInt := result.(int)
	fmt.Println(resultInt)

	// TYPE SWITCH - CARA AMAN UNTUK MENGECEK TYPE:
	// Type switch menggunakan keyword 'type' dalam switch statement
	// Ini adalah cara yang aman untuk menangani berbagai tipe data
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)   
	default:
		fmt.Println("Unknown type") 
	}
}