package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type File struct {
	FileName string
	Lines    int
	Words    int
	Error    error
}

func processFile(filename string) (int, int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	wordCount := 0

	for scanner.Scan() {
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

	for i, fileName := range files{
		content := fmt.Sprintf("This is file %d\nIt has multiple lines\nAnd words to count", i+1)

		os.WriteFile(fileName, []byte(content), 0644)

		defer os.Remove(fileName)
	}

	start := time.Now()
	results := make(chan File, len(files))

	for _, fileName := range files{
		go func(f string){
			lines, words, err := processFile(f)
			results <- File{
				FileName: f,
				Lines: lines,
				Words: words,
				Error: err,
			}
		}(fileName)

	}

	totalLines, totalWords := 0,0

	for i :=0; i < len(files); i++{
		result := <-results

		if result.Error != nil{
			fmt.Printf("Error processing %s: %v\n", result.FileName, result.Error)
		}else{
			fmt.Printf("✅ %s - Lines: %d, Words: %d\n", result.FileName, result.Lines, result.Words)
            totalLines += result.Lines
            totalWords += result.Words
		}
	}

	fmt.Printf("Total Lines: %d, Words: %d\n", totalLines, totalWords)
    fmt.Printf("Time taken: %v\n", time.Since(start))

}