package main

import "fmt"

func main() {
	var names = [3]string{
		"john",
		"doe",
		"wick",
	}

	fmt.Println(names)
	fmt.Println(names[0], names[1], names[2])
	fmt.Println(len(names))

	var values = [...]int{
		90,
		80,
		95,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}
