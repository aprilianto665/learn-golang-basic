package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	// request.Header.Get() - mengambil nilai header dari request client
	// Header names case-insensitive: "content-type" = "Content-Type"
	contentType := request.Header.Get("content-type")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	// request.Header.Add() - menambahkan header ke request untuk testing
	// Content-Type memberitahu server format data yang dikirim
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body)) // Output: "application/json"
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	// writer.Header().Add() - menambahkan header ke response
	// X-Powered-By adalah custom header untuk identifikasi server
	// PENTING: Header harus ditambahkan SEBELUM menulis body
	writer.Header().Add("X-Powered-By", "Golang")
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T){
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body)) // Output: "OK"

	// response.Header().Get() - mengambil header dari response yang direkam
	fmt.Println(response.Header.Get("X-Powered-By")) // Output: "Golang"
}