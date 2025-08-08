package main

import "fmt"

// Custom error type untuk validation error
// Struct yang mengimplementasikan interface error
type validationError struct {
	Message string
}

// Method Error() untuk mengimplementasikan interface error
// Setiap type yang memiliki method Error() string adalah error
func (e *validationError) Error() string {
	return e.Message
}

type notFoundError struct {
	Message string
}

func (e *notFoundError) Error() string {
	return e.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}
	if id != "john" {
		return &notFoundError{"not found error"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	
	if err != nil {
		// Type assertion menggunakan switch statement (lebih clean dari if-else)
		// Pattern: switch variable := interface.(type)
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Not Found Error:", finalError.Error())
		default:
			fmt.Println("Unknown Error:", finalError.Error())
		}
	} else {
		fmt.Println("Success")
	}
}