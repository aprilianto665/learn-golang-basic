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

// Function dengan variadic parameters
// Dapat menerima jumlah parameter yang tidak terbatas
// Parameter variadic harus diletakkan di akhir
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
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

	// Memanggil function dengan variadic parameters
	// Dapat mengirimkan jumlah parameter yang tidak terbatas
	total := sumAll(10,10,10,10,10)
	fmt.Println(total)

	// Dapat mengirimkan slice dengan operator ...
	// Operator ... akan mengubah slice menjadi variadic parameters
	// Jadi, slice harus memiliki tipe data yang sama dengan parameter variadic
	numbers := []int{10,10,10}
	total2 := sumAll(numbers...)
	fmt.Println(total2)
}