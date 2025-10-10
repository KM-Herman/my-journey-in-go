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

	var wg sync.WaitGroup
	totalLines, totalWords := 0,0
	var mu sync.Mutex

	for _, fileName := range files{
		wg.Add(1)
		go func(f string){
			defer wg.Done()

			lines, words, err := processFile(f)
		}
	}

}