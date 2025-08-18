package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// LogJson - Fungsi helper untuk mengkonversi data ke format JSON dan menampilkannya
// Parameter: data interface{} - menerima tipe data apapun
// Proses: menggunakan json.Marshal() untuk encoding ke JSON bytes, lalu convert ke string
func LogJson(data interface{}){
	bytes, err := json.Marshal(data) // Marshal mengkonversi Go data ke JSON format
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T){
	LogJson("Rusdi")
	LogJson(123)
	LogJson(true)
	LogJson([]string{"Rusdi", "Amba", "Fuad"})
}