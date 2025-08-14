package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	
	// new(http.Cookie) - membuat instance cookie baru
	cookie := new(http.Cookie)
	cookie.Name = "X-AMB-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"                                 // Path dimana cookie berlaku

	// http.SetCookie() - mengirim cookie ke client melalui Set-Cookie header
	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success create cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {

	// request.Cookie() - mengambil cookie berdasarkan nama
	cookie, err := request.Cookie("X-AMB-Name")
	if err != nil {
		fmt.Fprint(writer, "No cookie")
	} else {
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=John", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	// recorder.Result().Cookies() - mengambil semua cookie dari response
	cookies := recorder.Result().Cookies()
	if len(cookies) == 0 {
		t.Error("No cookies set")
	} else {

		// Menampilkan nama dan nilai cookie yang di-set
		for _, cookie := range cookies {
			fmt.Println(cookie.Name, cookie.Value) // Output: X-AMB-Name John
		}
	}
}

func TestGetCookie(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/get-cookie", nil)

	// &http.Cookie{} - menggunakan pointer karena AddCookie() membutuhkan *http.Cookie
	// Pointer lebih efisien untuk struct besar dan konsisten dengan API Go
	cookie := &http.Cookie{Name: "X-AMB-Name", Value: "John"}

	// request.AddCookie() - menambahkan cookie ke request untuk testing
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body)) // Output: "Hello John"
}