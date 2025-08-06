package main

import "fmt"

// Type alias untuk function - membuat kode lebih readable
type Filter func(string) string

// Higher order function - menerima function sebagai parameter
func sayHelloWithFilter(name string, filter Filter) string {
	return "Hello " + filter(name)
}

// Function sebagai parameter
func spamFilter(name string) string{
	if name == "Anjing"{
		return "..."
	}
	return name
}

// Type alias untuk function dengan return bool
type BlackList func(string) bool

// Function yang menerima function sebagai parameter
func registerUser(name string, blacklist BlackList){
	if blacklist(name){
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main(){
	// Function sebagai parameter
	helloFilter := sayHelloWithFilter("Anjing", spamFilter)
	fmt.Println(helloFilter)

	// Anonymous function - function tanpa nama langsung sebagai parameter
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}