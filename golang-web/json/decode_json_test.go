package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T){
	// JSON string yang akan di-decode (menggunakan backtick untuk raw string)
	jsonString := `{"FirstName":"Pharell","MiddleName":"Brown","LastName":"Williams","Age":30,"Married":true}`
	
	// Convert string ke []byte karena json.Unmarshal membutuhkan []byte
	jsonBytes := []byte(jsonString)

	// Buat pointer ke struct Customer kosong untuk menampung hasil decode
	// PENTING: Harus pointer (&Customer{}) bukan value (Customer{})
	// Karena json.Unmarshal perlu memodifikasi struct yang sudah ada
	// Jika pakai value, Go akan pass by copy dan perubahan tidak tersimpan
	customer := &Customer{}

	// Unmarshal JSON bytes ke struct Customer
	// Parameter kedua harus pointer agar bisa dimodifikasi
	// JSON keys akan dipetakan ke field struct yang sesuai
	json.Unmarshal(jsonBytes, customer)

	// Print seluruh struct - format: &{Pharell Brown Williams 30 true}
	fmt.Println(customer)
	
	// Print individual field yang sudah ter-decode
	fmt.Println(customer.FirstName)  // Output: Pharell
	fmt.Println(customer.MiddleName) // Output: Brown
	fmt.Println(customer.LastName)   // Output: Williams
	fmt.Println(customer.Age)        // Output: 30
	fmt.Println(customer.Married)    // Output: true
}