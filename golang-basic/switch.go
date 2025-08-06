package main

import "fmt"

func main() {
	grade := "A"

	// Switch statement dasar
	// Tidak perlu break seperti di bahasa lain
	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good")
	case "C":
		fmt.Println("Fair")
	default:
		fmt.Println("Poor")
	}

	name := "John Doe"

	// Switch dengan short variable declaration
	// Variable 'length' hanya tersedia dalam scope switch ini
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	score := 85

	// Switch tanpa expression (seperti if-else chain)
	// Berguna untuk multiple kondisi yang berbeda
	switch {
	case score > 80:
		fmt.Println("Excellent")
	case score > 60:
		fmt.Println("Good")
	default:
		fmt.Println("Poor")
	}
}
