package main

import (
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T){
	// http.Dir() - mengkonversi string path ke tipe http.Dir
	// "./resources" adalah folder yang berisi file static
	directory := http.Dir("./resources")
	// http.FileServer() - handler untuk serve files dari directory
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	// http.StripPrefix() - menghapus prefix "/static" dari URL sebelum mencari file
	// Contoh: /static/style.css → mencari file ./resources/style.css
	// mux.Handle() digunakan untuk handler, bukan HandleFunc untuk function
	mux.Handle("/static/",http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}