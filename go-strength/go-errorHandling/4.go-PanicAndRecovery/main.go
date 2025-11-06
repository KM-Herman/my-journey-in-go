package main

import (
	"fmt"
	"log"
)


func safeDivide(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered from panic: %v\n", r)
            fmt.Println("Division failed, but program continues")
        }
    }()
    
    result := divideWithPanic(a, b)
    fmt.Printf("Division result: %d\n", result)
}

func divideWithPanic(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}

func runSafeGoroutine() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Goroutine panic recovered: %v", r)
        }
    }()
    
    go func() {
        fmt.Println("Goroutine started")
        panic("something went wrong in goroutine")
    }()
    
}

func processRequest(request string) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Request %s caused panic: %v\n", request, r)
            // Log the error
            log.Printf("Panic recovered for request: %s", request)
        }
    }()
    
    handleRequest(request)
}

func handleRequest(request string) {
    if request == "panic request" {
        panic("deliberate panic for testing")
    }
    fmt.Printf("Successfully processed: %s\n", request)

}

func main() {
    
    // Basic panic recovery
    safeDivide(10, 0)
    safeDivide(20, 5)
    
    // Recovery in goroutines
    runSafeGoroutine()
    
    // Web server like panic recovery
    processRequest("valid request")
    processRequest("panic request")
    
    fmt.Println("Program continues normally...")
}