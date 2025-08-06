package main

import "fmt"

func main() {
	counter := 1

	// For loop seperti while loop
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// For loop dengan init, condition, dan post statement
	for count := 1; count <= 10; count++ {
		fmt.Println("Perulangan ke ", count)
	}

	names := []string{"john", "doe", "wick"}

	// For loop dengan index untuk slice/array
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// For range - mendapat index dan value
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	// For range - hanya value (ignore index dengan _)
	for _, name := range names {
		fmt.Println(name)
	}

	person := map[string]string{
		"name": "John Doe",
		"age":  "30",
	}

	// For range untuk map - mendapat key dan value
	for key, value := range person {
		fmt.Println("Key", key, "=", value)
	}
}