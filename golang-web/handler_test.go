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