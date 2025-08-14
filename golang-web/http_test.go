package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest" // Package untuk testing HTTP handler tanpa server nyata
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

// TestHttp - unit test untuk HTTP handler tanpa menjalankan server
func TestHttp(t *testing.T){
	// httptest.NewRequest - membuat HTTP request palsu untuk testing
	// Parameter: method, URL, body (nil = tidak ada body)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	// httptest.NewRecorder - merekam response dari handler untuk diverifikasi
	recorder := httptest.NewRecorder()

	// Memanggil handler dengan request palsu dan recorder
	HelloHandler(recorder, request)

	// Mengambil hasil response dari recorder
	response := recorder.Result()
	// Membaca body response dan mengkonversi ke string
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body)) // Output: "Hello World"
}