package main

import (
	"fmt"
	"net/http"
	"testing"
)

// RedirectTo - handler tujuan redirect
func RedirectTo(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Hello Redirect")
}

// RedirectFrom - handler yang melakukan redirect ke URL lain
func RedirectFrom(writer http.ResponseWriter, request *http.Request){
	// http.Redirect() - mengirim HTTP redirect response ke client
	// Parameter: writer, request, target URL, status code
	// http.StatusTemporaryRedirect = 307 (Temporary Redirect)
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

// RedirectOut - handler yang melakukan redirect ke URL eksternal
func RedirectOut(writer http.ResponseWriter, request *http.Request){
	http.Redirect(writer, request, "https://www.google.com", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}