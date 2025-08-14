package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		// writer.WriteHeader() - mengatur HTTP status code response
		// http.StatusBadRequest = 400 (Bad Request)
		// PENTING: WriteHeader harus dipanggil SEBELUM menulis body
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Name is empty")
	} else {
		// Jika tidak ada WriteHeader(), default status adalah 200 OK
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCode(t *testing.T){
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	// response.StatusCode - angka status code (400)
	fmt.Println(response.StatusCode)
	// response.Status - string lengkap status ("400 Bad Request")
	fmt.Println(response.Status)
	fmt.Println(string(body)) // Output: "Name is empty"
}
