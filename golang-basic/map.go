package main

import "fmt"

func main(){
	person := map[string]string{}
	person["name"] = "John Doe"
	person["age"] = "30"

	fmt.Println(person)

	product := map[string]string{
		"name":  "Laptop",
		"price": "1000 USD",
	}

	fmt.Println(product["name"])
	fmt.Println(product["price"])

	delete(person, "age")
	fmt.Println(person)
}