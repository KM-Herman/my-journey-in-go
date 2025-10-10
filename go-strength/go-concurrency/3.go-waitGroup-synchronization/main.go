package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

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

	//Without waitGroup your program could exit before goroutines complete
	var wg sync.WaitGroup // waitGroup is used to wait for a collection of goroutines to finish
	totalLines, totalWords := 0,0
	var mu sync.Mutex // mutex ensures that only one goroutine at a time can access a critical section of code


	for _, fileName := range files{
		wg.Add(1)
		go func(f string){
			defer wg.Done()

			lines, words, err := processFile(f)

			if err != nil {
				fmt.Printf("Error processing %s: %v\n", f, err)
                return
			}

			fmt.Printf("%s - Lines: %d, Words: %d\n", f, lines, words)

			// multiple goroutines might try to update them at the same time so
			// without a lock you could get a race condition
			// only one goroutine can enter the locked section at a time

			mu.Lock()
			totalLines += lines
			totalWords += words
			mu.Unlock()
		}(fileName)
	}

	wg.Wait() // ensures main doesn’t exit early without finishing ll goroutines

	fmt.Printf("Total Lines: %d, Words: %d\n", totalLines, totalWords)
    fmt.Printf("Time taken: %v\n", time.Since(start))
}