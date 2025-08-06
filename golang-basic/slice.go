package main

import "fmt"

func main() {
	// Membuat array dengan ukuran tetap
	names := [...]string{"john", "doe", "wick"," smith", "jane"}
	
	// Slice adalah potongan dari array
	// Syntax: array[start:end] - end tidak termasuk
	slice := names[1:4]  // index 1,2,3 -> "doe", "wick", " smith"
	slice1 := names[2:]  // dari index 2 sampai akhir
	slice2 := names[:3]  // dari awal sampai index 2
	slice3 := names[:]   // seluruh array

	fmt.Println(slice)
	fmt.Println("length: ",len(slice))   // panjang slice
	fmt.Println("capacity: ",cap(slice)) // kapasitas slice
	fmt.Println(slice1, slice2, slice3)

	// PENTING: Slice adalah reference ke array asli
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	daySlice1 := days[5:]  // slice dari index 5 ke akhir
	daySlice1[0] = "holiday"  // mengubah "saturday" menjadi "holiday"
	daySlice1[1] = "holiday"  // mengubah "sunday" menjadi "holiday"
	fmt.Println(days) // array asli ikut berubah karena slice adalah pointer

	// append() membuat array baru jika kapasitas tidak cukup
	daySlice2 := append(daySlice1, "new holiday")
	daySlice2[0] = "new holiday"  // mengubah elemen pertama
	fmt.Println(daySlice2) 
	fmt.Println(days) // array asli tidak berubah karena daySlice2 sudah menunjuk ke array baru

	// Membuat slice dengan make(type, length, capacity)
	newSlice := make([]string, 2, 5)  // length=2, capacity=5
	newSlice[0] = "john"
	newSlice[1] = "doe"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))  // length = 2
	fmt.Println(cap(newSlice))  // capacity = 5

	// append() menambah elemen ke slice
	newSlice2 := append(newSlice, "wick")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))  // length bertambah menjadi 3
	fmt.Println(cap(newSlice2))  // capacity tetap 5

	// Menyalin slice dengan copy()
	fromSlice := days[:]  // slice dari seluruh array days
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)  // menyalin isi fromSlice ke toSlice
	
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Perbedaan Array vs Slice
	thisIsArray := [...]int{1,2,3,4,5}  // Array: ukuran tetap [...]
	thisIsSlice := []int{1,2,3,4,5}     // Slice: ukuran dinamis []

	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)
}