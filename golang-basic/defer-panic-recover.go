package main

import "fmt"

// endApp adalah function yang akan dipanggil saat defer dieksekusi
func endApp() {
	fmt.Println("Ending the application...")
	// recover() digunakan untuk menangkap panic dan mengembalikan nilai yang dipanic
	// recover() hanya berfungsi di dalam deferred function
	message := recover()
	fmt.Println("Error:", message)
}

// runApp mendemonstrasikan penggunaan defer dan panic
func runApp(error bool){
	// defer: menunda eksekusi function hingga function saat ini selesai
	// endApp() akan dipanggil terakhir, bahkan jika terjadi panic
	defer endApp()

	if error {
		// panic: menghentikan eksekusi normal dan memicu panic state
		// Semua deferred function akan tetap dieksekusi sebelum program crash
		panic("Error!")
	}
}

func main() {
	// Memanggil runApp dengan parameter true untuk memicu panic
	// Output: "Ending the application..." kemudian "Error: Error!"
	runApp(true)
}
