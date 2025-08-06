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

// Multiple return values
func getFullName() (string,string){
	return "John", "Doe"
}

// Named return values - variabel sudah dideklarasi di signature
func getCompleteName() (firstName3, middleName3,lastName3 string){
	firstName3 = "John"
	middleName3 = "Wick"
	lastName3 = "Doe"
	return firstName3, middleName3,lastName3
}

// Variadic parameters - menerima parameter tidak terbatas
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Function sebagai first class citizen - function dapat disimpan kedalam variable
func getGoodbye(name string) string {
	return "Good bye " + name
}

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

// Recursive function - function yang memanggil dirinya sendiri
func sumRecursive(a, b int) int {
	if b == 0 {
		return a
	} else {
		return sumRecursive(a+1,b-1)
	}
}

func main(){
	// Memanggil function tanpa return
	sayHelloTo("John", "Doe")

	// Menyimpan return value ke variable
	hello := getHello("John")
	fmt.Println(hello)

	// Multiple return values
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// Mengabaikan return value dengan _
	firstName2, _ := getFullName()
	fmt.Println(firstName2)

	// Named return values
	firstName3, middleName3, lastName3 := getCompleteName()
	fmt.Println(firstName3, middleName3,lastName3)

	// Variadic parameters
	total := sumAll(10,10,10,10,10)
	fmt.Println(total)

	// Slice ke variadic dengan operator ...
	numbers := []int{10,10,10}
	total2 := sumAll(numbers...)
	fmt.Println(total2)

	// Function disimpan ke variable
	goodbye := getGoodbye
	fmt.Println(goodbye("John"))

	// Function sebagai parameter
	helloFilter := sayHelloWithFilter("Anjing", spamFilter)
	fmt.Println(helloFilter)

	// Anonymous function - function tanpa nama langsung sebagai parameter
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})

	// Memanggil recursive function
	fmt.Println(sumRecursive(5,5))
}