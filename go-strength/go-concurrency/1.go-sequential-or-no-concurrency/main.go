package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func processFile(filename string) (int, int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, 0, err
	}

	defer file.Close() //defer statement schedules a function call to be executed after the surrounding function returns

	scanner := bufio.NewScanner(file)
	lineCount := 0
	wordCount := 0

	for scanner.Scan(){
		lineCount++
		words := strings.Fields(scanner.Text())
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	time.Sleep(100 * time.Microsecond)

	return lineCount, wordCount, nil
}

func main(){
	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt"}

	for i, fileName := range files {
		content := fmt.Sprintf("This is file %d\nIt has multiple lines\nAnd words to count", i+1)

		os.WriteFile(fileName, []byte(content), 0644)

		defer os.Remove(fileName)
	}

	start := time.Now()
	totalLines, totalWords := 0,0


	for _, fileName := range files {
		lines, words, err := processFile(fileName)

		if err != nil{
			fmt.Printf("Error processing %s: %v\n", fileName, err)
            continue
		}

		fmt.Printf("%s - Lines: %d, Words: %d\n", fileName, lines, words)
        totalLines += lines
        totalWords += words
	}

	fmt.Printf("Total - Lines: %d, Words: %d\n", totalLines, totalWords)
	fmt.Printf("Time taken: %v\n", time.Since(start))

}