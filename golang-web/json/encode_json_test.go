package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(data interface{}){
	bytes, err := json.Marshal(data)
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