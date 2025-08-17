package main

import (
	"fmt"
	"net/http"
	"testing"
)

// LogMiddleware - struct yang mengimplementasikan http.Handler interface
type LogMiddleware struct {
	Handler http.Handler // Handler yang akan di-wrap oleh middleware
}

// ServeHTTP - implementasi http.Handler interface untuk middleware
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Kode yang dijalankan SEBELUM handler utama
	fmt.Println("Before execute handler")
	// Memanggil handler yang di-wrap (mux dalam hal ini)
	middleware.Handler.ServeHTTP(writer, request)
	// Kode yang dijalankan SETELAH handler utama selesai
	fmt.Println("After execute handler")
}

// TestMiddleware - test middleware pattern dengan logging
func TestMiddleware(t *testing.T){
	// Membuat handler utama (mux)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler executed")
		fmt.Fprintf(writer, "Hello Middleware")
	})
	// Membuat instance middleware dan wrap handler utama
	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux // Mux di-wrap oleh middleware

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: logMiddleware, // Server menggunakan middleware sebagai handler utama
		}

	// Flow: Request → LogMiddleware → Mux → Response
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}