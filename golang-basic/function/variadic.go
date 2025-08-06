package main

import "fmt"

// Variadic parameters - menerima parameter tidak terbatas
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main(){
	// Variadic parameters
	total := sumAll(10,10,10,10,10)
	fmt.Println(total)

	// Slice ke variadic dengan operator ...
	numbers := []int{10,10,10}
	total2 := sumAll(numbers...)
	fmt.Println(total2)
}