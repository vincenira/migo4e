/*
Try to implement a concurrent version of wc(1) that uses shared memory.
We will implement it using the mutex to lock the variable to share.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"
)

var (
	readString []string
	mutexRead  sync.RWMutex
	wg         sync.WaitGroup
	totalLines int
	totalWords int
	totalChars int
)

func readFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)

	mutexRead.Lock()
	defer mutexRead.Unlock()

	for {
		// ReadString() returns two values: the string that was read and an error variable.
		line, err := r.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				readString = append(readString, line)
			}
			break
		}

		if err != nil {
			fmt.Printf("error reading file %s", err)
		}
		readString = append(readString, line)
	}
}

func wordByWord() {
	defer wg.Done()
	defer mutexRead.RUnlock()

	total := 0
	re := regexp.MustCompile("[^\\s]+")
	mutexRead.RLock()
	for _, line := range readString {
		if len(line) != 0 {
			words := re.FindAllString(line, -1)
			total += len(words)
		}
	}
	totalWords = total
}

func lineByLine() {
	defer wg.Done()
	defer mutexRead.RUnlock()
	mutexRead.RLock()
	totalLines = len(readString)
}

func charByChar() {
	defer wg.Done()
	defer mutexRead.RUnlock()
	total := 0
	totalChars = 0
	mutexRead.RLock()
	for _, line := range readString {
		total += len(string(line))
	}
	totalChars = total
}

func printTotalResult() {
	defer wg.Done()
	defer mutexRead.RUnlock()
	mutexRead.RLock()
	if totalLines != 0 {
		fmt.Printf("%d  ", totalLines)
	}
	if totalWords != 0 {
		fmt.Printf("%d ", totalWords)
	}
	if totalChars != 0 {
		fmt.Printf("%d ", totalChars)
	}
	fmt.Printf("total\n")
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: wcShaMem <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		readFile(file)
		wg.Add(3)
		go lineByLine()
		go wordByWord()
		go charByChar()
		wg.Wait()
		wg.Add(1) // Add another goroutine to print the total result after all reading goroutines have finished.
		go printTotalResult()
		wg.Wait()
	}
}
