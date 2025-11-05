package main

import (
	"fmt"
	"time"
)

type ValidatorError struct {
	Field string
	Message string
	value interface{}
}

type NetworkError struct {
	URL string
	StatusCode int
	Timeout time.Duration
}

func (e ValidatorError) Error() string {
	return fmt.Sprintf("Validation error on field '%s': %s (got %v)", e.Field, e.Message, e.value)
}

func (e NetworkError) Error() string {
	if e.Timeout > 0 {
        return fmt.Sprintf("network timeout to %s after %v", e.URL, e.Timeout)
    }
    return fmt.Sprintf("network error to %s: status %d", e.URL, e.StatusCode)
}

func ValidateUser(fullName string, age int) error {
	if fullName == "" {
		return ValidatorError{
			Field: "FullName",
			Message: "Full name cannot be empty",
			value: fullName,
		}
	}

	if age < 0 || age > 120 {
		return ValidatorError{
			Field: "Age",
			Message: "Age must be between 0 and 120",
			value: age,
		}
	}

	return nil
}

func MakeAPIRequest(url string) error {
	return NetworkError{
		URL: url,
		Timeout: 30 * time.Second,
	}
}


func main(){

	if err := ValidateUser("", -25); err != nil {
		fmt.Printf("Validation failed: %v\n", err)

			if valEr, ok := err.(ValidatorError); ok {
				fmt.Printf("Field: %s, Message: %s, Value: %v\n", valEr.Field, valEr.Message, valEr.value)
			}
	}

	if err:= MakeAPIRequest("https://api.github.com/users"); err != nil {
		fmt.Printf("API request failed: %v\n", err)

		switch e := err.(type) {
			case NetworkError:
				if e.Timeout > 0 {
					fmt.Println("this is a timeout error")
				} else {
					fmt.Printf("HTTP status code: %d\n", e.StatusCode)
				}
				default:
					fmt.Println("unknown error type")
		}
	}
}
