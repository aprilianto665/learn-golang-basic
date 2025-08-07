package main

import "fmt"

// Struct adalah tipe data yang mengelompokkan beberapa field dengan tipe berbeda
// Customer adalah custom type yang memiliki 3 field
type Customer struct {
	Name, Address string // Deklarasi multiple field dengan tipe yang sama
	Age           int    // Field dengan tipe berbeda
}

func main() {
	// Cara 1: Deklarasi struct kosong kemudian assign nilai per field
	var customer Customer
	customer.Name = "Alice"
	customer.Address = "789 Pine Street"
	customer.Age = 28
	fmt.Println(customer) // Output: {Alice 789 Pine Street 28}

	// Cara 2: Struct literal dengan nama field (recommended)
	// Lebih readable dan tidak bergantung pada urutan field
	john := Customer{
		Name:    "John Doe",
		Address: "123 Elm Street",
		Age:     30,
	}
	fmt.Println(john) // Output: {John Doe 123 Elm Street 30}

	// Cara 3: Struct literal tanpa nama field (harus sesuai urutan)
	// Kurang readable, tidak disarankan untuk struct dengan banyak field
	jane := Customer{"Jane Doe", "456 Oak Avenue", 25}
	fmt.Println(jane) // Output: {Jane Doe 456 Oak Avenue 25}
}
