package main

import "fmt"

// Multiple return values
func getFullName() (string,string){
	return "John", "Doe"
}

// Named return values - variabel sudah dideklarasi di signature
func getCompleteName() (firstName, middleName, lastName string){
	firstName = "John"
	middleName = "Wick"
	lastName = "Doe"
	return firstName, middleName, lastName
}

func main(){
	// Multiple return values
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// Mengabaikan return value dengan _
	firstName2, _ := getFullName()
	fmt.Println(firstName2)

	// Named return values
	firstName3, middleName3, lastName3 := getCompleteName()
	fmt.Println(firstName3, middleName3, lastName3)
}