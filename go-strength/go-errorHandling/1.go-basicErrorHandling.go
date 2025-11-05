package main

import(
	"fmt"
	"os"
	"strconv"
)

func divide(num1, num2 int) (int, error){
	if (num2 == 0){
		return 0, fmt.Errorf("division by zero is not allowed")
	}else{
		return num1/num2, nil
	}
}

func stringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to read file %s: %w", filename, err)
	}
	return string(data), nil
}

func main() {

	//1. Division by zero error handling

	result, err := divide(10, 0)
	if err != nil{
		fmt.Printf("Error: %v\n", err)
	}else{
		fmt.Printf("Result: %v\n", result)
	}

	//2. String to integer conversion error handling

	value, err := stringToInt("123")
    if err != nil {
        fmt.Printf("Conversion error: %v\n", err)
    } else {
        fmt.Printf("Converted value: %d\n", value)
    }

   //3. File reading error handling

	content, err := readFile("textFile.txt")
    if err != nil {
        fmt.Printf("File error: %v\n", err)
    } else {
        fmt.Printf("File content: %s\n", content)
    }
}