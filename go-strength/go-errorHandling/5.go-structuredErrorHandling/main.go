package main

import (
	"encoding/json" 
	"errors"       
	"fmt"
	"net/http"
	"time"
)

type AppError struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Details string    `json:"details,omitempty"`
	Time    time.Time `json:"timestamp"`
	TraceID string    `json:"trace_id,omitempty"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("Code:%d\n Message: %s\n Details: %s", e.Code, e.Message, e.Details)
}

func NewAppError(code int, message, details string) AppError {
	return AppError{
		Code:    code,
		Message: message,
		Details: details,
		Time:    time.Now(),
	}
}

type ValidationError struct{ s string }

func (e ValidationError) Error() string { return e.s }

func NewValidationError(msg string) ValidationError { 
	return ValidationError{msg}
}

type ErrorHandler struct {
	Logger func(error)
}

func (h *ErrorHandler) Handle(err error) {
	if h.Logger != nil {
		h.Logger(err)
	}
	appErr := h.toAppError(err)
	fmt.Printf("Handled error occurred: %s\n", appErr)
}

func (h *ErrorHandler) toAppError(err error) AppError {
	switch e := err.(type) {
	case AppError:
		return e
	case ValidationError:
		return NewAppError(http.StatusBadRequest, "Validation Error", e.Error())
	default:
		return NewAppError(http.StatusInternalServerError, "Unknown Error", err.Error())
	}
}


func validateUser(username string, age int) error {
    if username == "" {
        return NewValidationError("Username cannot be empty") 
    }
    if age < 18 {
        return NewValidationError("User must be 18 or older")
    }
    return nil
}

func processPayment(amount float64) error {
	if amount <= 0 {
		return NewAppError(
			http.StatusBadRequest,
			"Invalid payment amount",
			fmt.Sprintf("Amount must be positive, got %.2f", amount),
		)
	}

	if amount > 1000 {
		return NewAppError(
			http.StatusPaymentRequired,
			"Insufficient funds",
			"Transaction amount exceeds account balance",
		)
	}

	return nil
}

func fetchFromDatabase() error {
	return NewAppError(
		http.StatusInternalServerError,
		"Database connection failed",
		"Unable to connect to primary database replica",
	)
}

func processUserRequest() error {
	return NewAppError(
		429, 
		"Rate limit exceeded", 
		"Too many requests, please try again later",
	)
}

func sendErrorResponse(err error) {
	var appErr AppError
	if errors.As(err, &appErr) {
		jsonData, err := json.MarshalIndent(appErr, "", " ")
        if err != nil {
            fmt.Printf("\nError marshaling JSON: %v\n", err)
            return
        }
		fmt.Printf("\nHTTP Error Response (JSON):\n%s\n", string(jsonData))
	} else {
        fmt.Printf("\nHTTP Error Response: Generic error (%v)\n", err)
    }
}

func handleHTTPRequest(handler *ErrorHandler) {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("panic recovered: %v", r)
			handler.Handle(err)
            sendErrorResponse(err) 
		}
	}()

	fmt.Println("\nSimulating HTTP Request Handling")
	err := processUserRequest()
	if err != nil {
		handler.Handle(err)
		sendErrorResponse(err) // Send the standardized JSON back to the client
	}
}

func main() {

	handler := &ErrorHandler{
		Logger: func(err error) {
			fmt.Printf("Logged error: %v\n", err)
		},
	}

    fmt.Println("\nScenario 1: Validation Error")
	handler.Handle(validateUser("", 25))
    
    fmt.Println("\nScenario 2: Payment Error")
	handler.Handle(processPayment(0))

    fmt.Println("\nScenario 3: Database Error")
	handler.Handle(fetchFromDatabase())

	handleHTTPRequest(handler)
}