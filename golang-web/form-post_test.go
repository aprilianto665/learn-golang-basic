package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	// request.ParseForm() - parsing form data dari request body
	// Harus dipanggil sebelum mengakses PostForm atau Form
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	// request.PostForm - hanya berisi data dari request body (POST data)
	// Berbeda dengan request.Form yang gabungan URL query + POST data
	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T){
	// strings.NewReader - konversi string ke io.Reader (request body butuh stream, bukan string)
	// Format: key1=value1&key2=value2 (URL encoded)
	requestBody := strings.NewReader("first_name=John&last_name=Doe")
	request := httptest.NewRequest("POST", "http://localhost:8080/", requestBody)
	
	// Content-Type penting: memberitahu server format data yang dikirim
	// application/x-www-form-urlencoded = format form HTML standard
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body)) // Output: "Hello John Doe"
}

func FormPostValue(writer http.ResponseWriter, request *http.Request) {
	// request.PostFormValue() - otomatis memanggil ParseForm() jika belum dipanggil
	firstName := request.PostFormValue("first_name")
	lastName := request.PostFormValue("last_name")
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPostValue(t *testing.T){
	requestBody := strings.NewReader("first_name=Jane&last_name=Smith")
	request := httptest.NewRequest("POST", "http://localhost:8080/", requestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPostValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body)) // Output: "Hello Jane Smith"
}
