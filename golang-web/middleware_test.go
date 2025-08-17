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

// ServeHTTP - implementasi http.Handler interface untuk LogMiddleware
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Kode yang dijalankan SEBELUM handler utama
	fmt.Println("Before execute handler")
	// Memanggil handler yang di-wrap (mux dalam hal ini)
	middleware.Handler.ServeHTTP(writer, request)
	// Kode yang dijalankan SETELAH handler utama selesai
	fmt.Println("After execute handler")
}

// ErrorHandler - middleware untuk menangani panic dan error
type ErrorHandler struct {
	Handler http.Handler
}

// ServeHTTP - implementasi error handling middleware
func (middleware *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// defer recover() - menangkap panic yang terjadi di handler
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Internal error")
			// Mengirim status 500 dan pesan error ke client
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error: %s", err)
		}
	}()

	middleware.Handler.ServeHTTP(writer, request)
}

// TestMiddleware - test chaining multiple middleware
func TestMiddleware(t *testing.T){
	mux := http.NewServeMux()
	// Route normal yang tidak error
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler executed")
		fmt.Fprintf(writer, "Hello Middleware")
	})

	// Route yang akan panic untuk test error handling
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic executed")
		panic("Ups") // Simulasi error
	})

	// Chaining middleware: ErrorHandler → LogMiddleware → Mux
	logMiddleware := new(LogMiddleware)
	logMiddleware.Handler = mux // LogMiddleware wrap mux

	errorHandler := &ErrorHandler{
		Handler: logMiddleware, // ErrorHandler wrap LogMiddleware
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler, // Server menggunakan ErrorHandler sebagai entry point
		}

	// Flow: Request → ErrorHandler → LogMiddleware → Mux → Response
	// Akses /panic akan ditangkap oleh ErrorHandler
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}