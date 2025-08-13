package main

import (
	"fmt"
	"net/http"
	"testing"
)

// TestHandler - test untuk HTTP server dengan single handler
func TestHandler(t *testing.T){
	// HandlerFunc - fungsi yang menangani setiap HTTP request yang masuk
	var handler http.HandlerFunc = func(w http.ResponseWriter, request *http.Request){
		fmt.Fprint(w, "Hello World")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,       // Menggunakan handler custom untuk memproses semua request
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// TestServeMux - test untuk HTTP server dengan multiple routes (routing)
func TestServeMux(t *testing.T){
	// ServeMux - HTTP request multiplexer untuk routing ke handler yang berbeda
	// Berbeda dengan handler tunggal, mux bisa menangani banyak path/route
	mux := http.NewServeMux()

	// Route "/" - akan menangani semua request ke root dan path yang tidak terdaftar
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Hello World")
	})
	// Route "/hi" - hanya menangani request ke path /hi secara spesifik
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Hi")
	})

	// Route "/images/" - menangani semua request yang dimulai dengan /images/
	// Trailing slash penting: /images/photo.jpg akan ditangani oleh route ini
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Images")
	})

	// Route "/images/thumbnails/" - lebih spesifik dari /images/
	// ServeMux akan memilih route yang paling spesifik (longest match)
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Thumbnails")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux, // Menggunakan mux sebagai handler untuk routing
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// TestRequest - test untuk mengakses informasi dari HTTP request
func TestRequest(t *testing.T){
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request){
		// fmt.Fprintln(writer, data) - menulis data ke writer dengan newline di akhir
		// Parameter 1: writer (io.Writer) - tempat menulis output, di sini ke HTTP response
		// Parameter 2: data - nilai yang akan ditulis (request.Method = GET/POST/dll)
		fmt.Fprintln(writer, request.Method)
		// Sama seperti di atas, tapi menulis RequestURI (path + query string)
		// Contoh: /users?name=john akan menampilkan "/users?name=john"
		fmt.Fprintln(writer, request.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}