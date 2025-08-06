package main

import "fmt"

// Function sebagai first class citizen - function dapat disimpan kedalam variable
func getGoodbye(name string) string {
	return "Good bye " + name
}

func main(){
	// Function disimpan ke variable
	goodbye := getGoodbye
	fmt.Println(goodbye("John"))
}