package main

import "fmt"

func main(){
	// Membuat map kosong dan menambah data
	person := map[string]string{}
	person["name"] = "John Doe"
	person["age"] = "30"

	fmt.Println(person)

	// Membuat map dengan initial values
	product := map[string]string{
		"name":  "Laptop",
		"price": "1000 USD",
	}

	// Mengakses nilai dari map
	fmt.Println(product["name"])
	fmt.Println(product["price"])

	// Menghapus key dari map
	delete(person, "age")
	fmt.Println(person)
}