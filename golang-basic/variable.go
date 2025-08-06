package main

import "fmt"

func main() {
	// Short variable declaration dengan :=
	// Type otomatis ditentukan dari nilai
	hello := "Hello, World!"
	fmt.Println(hello)

	// Multiple variable declaration dengan var
	// Bisa mendeklarasikan beberapa variabel sekaligus
	var(
		name = "Ian"
		age = 18
	)

	fmt.Println(name, age)
}
