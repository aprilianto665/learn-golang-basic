package main

import "fmt"

// Function tanpa return
func sayHelloTo(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

// Function dengan return
func getHello(name string) string{
	return "Hello " + name
}

func main(){
	// Memanggil function tanpa return
	sayHelloTo("John", "Doe")

	// Menyimpan return value ke variable
	hello := getHello("John")
	fmt.Println(hello)
}