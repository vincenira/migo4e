/*
Try to implement a concurrent version of wc(1) that uses shared memory.
We will implement it using the mutex to lock the variable to share.
*/
package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"

	"golang.org/x/sync/semaphore"
)

var (
	readString []string
	totalLines int
	totalWords int
	totalChars int
	wg         sync.WaitGroup
)

// Maximum number of goroutines
var (
	Workers = 4
	sem     = semaphore.NewWeighted(int64(Workers))
)

func readFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)

	for {
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

func wordByWord(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer sem.Release(1)
	_ = sem.Acquire(ctx, 1)
	total := 0
	re := regexp.MustCompile("[^\\s]+")
	for _, line := range readString {
		if len(line) != 0 {
			words := re.FindAllString(line, -1)
			total += len(words)
		}
	}
	totalWords = total
}

func lineByLine(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer sem.Release(1)
	_ = sem.Acquire(ctx, 1)
	totalLines = len(readString)
}

func charByChar(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer sem.Release(1)
	_ = sem.Acquire(ctx, 1)
	total := 0
	totalChars = 0
	for _, line := range readString {
		total += len(string(line))
	}
	totalChars = total
}

func printTotalResult(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer sem.Release(1)
	_ = sem.Acquire(ctx, 1)
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
		fmt.Printf("usage: wcSemaphores <file1> [<file2> ...]\n")
		return
	}

	ctx := context.Background()
	for _, file := range args[1:] {
		readFile(file)
		wg.Add(3)
		go lineByLine(ctx, &wg)
		go wordByWord(ctx, &wg)
		go charByChar(ctx, &wg)
		wg.Wait()
		wg.Add(1)
		go printTotalResult(ctx, &wg)
		wg.Wait()
	}
}
