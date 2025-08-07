package main

import "fmt"

// Interface adalah kontrak yang mendefinisikan method apa saja yang harus dimiliki
// oleh sebuah type untuk mengimplementasikan interface tersebut
// PENTING: Nama method di interface HARUS sama persis dengan method di struct
type HasName interface {
	GetName() string    // Method untuk mendapatkan nama
	GetAddress() string // Method untuk mendapatkan alamat
}

// Function yang menerima parameter bertipe interface HasName
// Semua type yang mengimplementasikan interface HasName bisa digunakan di sini
func sayHello(value HasName){
	fmt.Println("Hello", value.GetName(), "from", value.GetAddress())
}

// Struct Person yang akan mengimplementasikan interface HasName
type Person struct {
	Name    string // Field untuk menyimpan nama
	Address string // Field untuk menyimpan alamat
}

// Method GetName() untuk struct Person
// Dengan adanya method ini, Person mengimplementasikan interface HasName
func (p Person) GetName() string {
	return p.Name
}

// Method GetAddress() untuk struct Person
// Method kedua yang diperlukan untuk mengimplementasikan interface HasName
func (p Person) GetAddress() string {
	return p.Address
}

// Function yang mendemonstrasikan penggunaan any dan interface{}
// any adalah alias untuk interface{} yang bisa menerima tipe data apapun
// interface{} adalah empty interface yang tidak memiliki method apapun
func Ups(value any) interface{} {
	// Mengecek nilai parameter menggunakan perbandingan langsung
	if value == 1 {
		return 1        // Mengembalikan integer
	} else if value == 2 {
		return true     // Mengembalikan boolean
	} else {
		return "Ups"    // Mengembalikan string
	}
}

func main(){
	// Membuat instance dari struct Person
	person := Person{Name: "John Doe", Address: "123 Main Street"}
	
	// Memanggil function sayHello dengan parameter person
	// Person bisa digunakan karena sudah mengimplementasikan interface HasName
	sayHello(person)

	// Memanggil function Ups dengan parameter string "test"
	// Karena parameter bukan 1 atau 2, akan mengembalikan "Ups"
	kosong := Ups("test")
	fmt.Println(kosong) // Output: Ups
}