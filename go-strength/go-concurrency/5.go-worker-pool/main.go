package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type FileResult struct {
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

func Worker(id int, jobs <-chan string, results chan<- FileResult, wg *sync.WaitGroup){
	defer wg.Done()

	for fileName := range jobs {
		fmt.Printf("Worker %d processing: %s\n", id, fileName)
		lines, words, err := processFile(fileName)
		results <-FileResult{
			FileName: fileName,
			Lines: lines,
			Words: words,
			Error: err,
		}
	}
}

func main(){
	files := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt", "file6.txt"}

	for i, fileName := range files {
		content := fmt.Sprintf("This is file %d\nIt has multiple lines\nAnd words to count", i+1)
        
		os.WriteFile(fileName, []byte(content), 0644)
        
		defer os.Remove(fileName)
    }

	start := time.Now()
	const numberWorkers = 3

	jobs := make(chan string, len(files))
	results := make(chan FileResult, len(files))

	//start workers

	var wg sync.WaitGroup
	for i := 0; i < numberWorkers; i++{
		wg.Add(1)
		go Worker(i, jobs, results, &wg)
	}

	// send jobs
	for _, fileName := range files {
		jobs <- fileName
	}
	close(jobs)

	//wait for completion

	wg.Wait()
	close(results)

	//process results

	totalLines, totalWords := 0,0

	for result := range results {
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

