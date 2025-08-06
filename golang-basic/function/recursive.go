package main

import "fmt"

// Recursive function - function yang memanggil dirinya sendiri
func sumRecursive(a, b int) int {
	if b == 0 {
		return a
	} else {
		return sumRecursive(a+1,b-1)
	}
}

func main(){
	// Memanggil recursive function
	fmt.Println(sumRecursive(5,5))
}