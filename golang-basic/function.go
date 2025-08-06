package main

import "fmt"

// Function tanpa return value
func sayHelloTo(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

// Function dengan return value
func getHello(name string) string{
	return "Hello " + name
}

// Function dengan multiple return values
func getFullName() (string,string){
	return "John", "Doe"
}

// Function dengan named return values
// Variable sudah dideklarasikan di signature function
func getCompleteName() (firstName3, middleName3,lastName3 string){
	firstName3 = "John"
	middleName3 = "Wick"
	lastName3 = "Doe"
	return firstName3, middleName3,lastName3
}

func main(){
	// Memanggil function tanpa return
	sayHelloTo("John", "Doe")

	// Memanggil function dengan return value
	hello := getHello("John")
	fmt.Println(hello)

	// Memanggil function dengan multiple return values
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// Mengabaikan salah satu return value dengan blank identifier (_)
	firstName2, _ := getFullName()
	fmt.Println(firstName2)

	// Memanggil function dengan named return values
	firstName3, middleName3, lastName3 := getCompleteName()
	fmt.Println(firstName3, middleName3,lastName3)
}