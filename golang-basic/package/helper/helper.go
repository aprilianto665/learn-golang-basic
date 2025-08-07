package helper

// Panduan Export/Import Function di Go:
// - Fungsi yang dimulai dengan huruf KAPITAL (SayHello) = PUBLIC/EXPORTED
//   Bisa diakses dari package lain
// - Fungsi yang dimulai dengan huruf kecil (sayHello) = PRIVATE
//   Hanya bisa diakses dalam package yang sama

func SayHello(name string) string {
	return "Hello " + name
}