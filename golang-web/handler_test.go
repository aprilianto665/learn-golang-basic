package main

import (
	"fmt"
	"net/http"
	"testing"
)

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