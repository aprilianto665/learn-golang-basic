package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Customer - Struct untuk merepresentasikan data customer
// Field dengan huruf kapital akan di-export dan bisa di-encode ke JSON
// Nama field akan menjadi key dalam JSON object
type Customer struct{
	FirstName string
	MiddleName string
	LastName string
	Age int
	Married bool
}

func TestJSONObject(t *testing.T){
	customer := Customer{
		FirstName: "Pharell",
		MiddleName: "Brown",
		LastName: "Williams",
		Age: 30,
		Married: true,
	}

	// Marshal struct ke JSON bytes
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}