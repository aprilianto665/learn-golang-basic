package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	// request.URL.Query() - mengambil semua query parameters dari URL
	// .Get("name") - mengambil nilai pertama dari parameter "name"
	// Jika parameter tidak ada, Get() mengembalikan string kosong ""
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQuery(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=John", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	// query tanpa .Get() mengembalikan url.Values (map[string][]string)
	query := request.URL.Query()
	// Akses langsung dengan key mengembalikan slice []string
	// names akan berisi ["John", "Doe"] dari ?name=John&name=Doe
	names := query["name"]
	// strings.Join() menggabungkan slice string dengan separator spasi
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T){
	// URL dengan parameter name yang sama dua kali: name=John&name=Doe
	// Ini akan menghasilkan slice ["John", "Doe"]
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=John&name=Doe", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}