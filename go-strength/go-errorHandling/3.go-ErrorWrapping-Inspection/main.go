package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
	Value   interface{}
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field '%s': %s (got %v)", 
        e.Field, e.Message, e.Value)
}

var(
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidAccess = errors.New("invalid access rights")
	ErrSystemBusy = errors.New("system is busy")
)

func ProcessUser(userID int) error {
	 if err := FindUser(userID); err != nil {
        return fmt.Errorf("failed to process user %d: %w", userID, err)
    }
    
    if err := ValidateData("user data"); err != nil {
        wrapped := fmt.Errorf("data validation failed: %w", err)
        return fmt.Errorf("user processing aborted: %w", wrapped)
    }

    return nil
}

func FindUser(userID int) error {
	 if userID == 0 {
        return fmt.Errorf("user lookup: %w", ErrUserNotFound)
    }
    if userID < 0 {
        return ValidationError{
            Field:   "userID",
            Message: "must be positive",
            Value:   userID,
        }
    }
    return nil
}

func ValidateData(data string) error {
    if strings.TrimSpace(data) == "" {
        return ValidationError{
            Field:   "data",
            Message: "cannot be empty",
            Value:   data,
        }
    }
    return nil
}

func ReadConfigFile(path string) error {
    _, err := os.Stat(path)
    if err != nil {
        return fmt.Errorf("config file check: %w", err)
    }
    
    data, err := os.ReadFile(path)
    if err != nil {
        return fmt.Errorf("read config %s: %w", path, err)
    }
    
    if len(data) == 0 {
        return fmt.Errorf("config file %s is empty", path)
    }
    
    return nil
}

func step1() error {
    if err := step2(); err != nil {
        return fmt.Errorf("step1: %w", err)
    }
    return nil
}

func step2() error {
    return fmt.Errorf("step2: %w", ErrSystemBusy)
}

func ComplexOperation() error {
    if err := step1(); err != nil {
        return fmt.Errorf("complex operation: %w", err)
    }
    return nil
}

func ErrorChain(err error) {
    fmt.Println("Error chain:")
    for err != nil {
        fmt.Printf("  - %v\n", err)
        err = errors.Unwrap(err)
    }
}

func main() {
    
    // Error wrapping demostration
    if err := ProcessUser(0); err != nil {
        fmt.Printf("Processing failed: %v\n", err)
        fmt.Printf("Unwrapped: %v\n", errors.Unwrap(err))
        
        if errors.Is(err, ErrUserNotFound) {
            fmt.Println("The user was not found in the system")
        }
        
        var ve *ValidationError
        if errors.As(err, &ve) {
            fmt.Printf("Validation error on field: %s\n", ve.Field)
        }
    }
    
    if err := ReadConfigFile("/app/config.json"); err != nil {
        fmt.Printf("Config error: %v\n", err)
        
        if errors.Is(err, fs.ErrNotExist) {
            fmt.Println("Config file does not exist")
        } else if errors.Is(err, fs.ErrPermission) {
            fmt.Println("Permission denied for config file")
        }
    }
    
    if err := ComplexOperation(); err != nil {
        fmt.Printf("Complex operation failed: %v\n", err)
        ErrorChain(err)
    }
}