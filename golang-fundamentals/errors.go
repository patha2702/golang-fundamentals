package main

import (
	"fmt"
	"errors"
)

// Error handling:
// - Go programs express errors with error values. 
// - An error is any type that implements the built-in error interface
// - type error interface {
//  	Error() string
//   }

// - When something can go wrong in a function, that function can return an error
//   as the last return value.
// - Any code that calls the function that can return error, should handle errors 
//   by testing whether the err is nil or not.

// Unlike programming lang like javascript, python, where we use try-catch block for 
// handling errors, in Go we explicitly deal with errors returned by functions using if-else.

func divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("Cannot divide %v by %v", num1, num2)
	}
	return num1/num2, nil
}


// Building custom error types
// - Since errors are just interfaces, we can create our own error types by 
//   implementing the error interface.

type userError struct{
	username string
}

func (e userError) Error() string {
	return fmt.Sprintf("%v user has a problem with their account.", e.username)
}

func getAccountStatus(username string) (string, error) {
	if username == "john" {
		return "", userError{username: username}
	}
	return "Success", nil
}


func main() {
	num1 := 20.0
	num2 := 0.0
	res, err := divide(num1, num2)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %w", err))
	} else {
		fmt.Printf("Result: %v", res)
	}

	// Custom error type
	status, err := getAccountStatus("johnson")
	fmt.Println(status, err)

	// Alternative way for creating custom error type
	// Using errors.New() function from errors package in Go standard library.
	var er error = errors.New("Something went wrong.")
	fmt.Println(er)
}
