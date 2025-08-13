package main

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	// ListenAndServe() memblokir eksekusi sampai server dihentikan
	err := server.ListenAndServe()
	if err != nil {
		// panic() menghentikan program karena server gagal start (misal: port sudah digunakan)
		panic(err)
	}
}