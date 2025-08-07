package main

import (
	"fmt"
	// Blank identifier (_) untuk import
	// Hanya menjalankan function init() dari package tersebut
	_ "golang-basic/function/internal"
)

var connection string

// Function init():
// - Dijalankan OTOMATIS sebelum function main()
// - Digunakan untuk inisialisasi awal (setup database, config, dll)
// - Tidak bisa dipanggil secara manual
// - Bisa ada lebih dari 1 function init() dalam 1 package
func init(){
	connection = "MYSQL"
	fmt.Println("Function init() dijalankan!")
}

func main(){
	fmt.Println("Function main() dijalankan!")
	fmt.Println("Connection:", connection)
}