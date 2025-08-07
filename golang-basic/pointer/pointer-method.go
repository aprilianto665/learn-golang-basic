package main

import "fmt"

type Man struct {
	Name string
}

// Method dengan pointer receiver (*Man)
// Menggunakan pointer agar bisa mengubah nilai asli struct
func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}

func main(){
	man := Man{"John"}
	man.Married()

	fmt.Println(man.Name)
}
