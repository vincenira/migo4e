/*
Try to implement a concurrent version of wc(1) that saves its output to a file.
Check ligne 360
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

/*
To do: adding the logic for channel for reading characters
*/
var (
	wg         sync.WaitGroup
	totalLines int
	totalWords int
	totalChars int
)

func readFile(path string, storeString chan []string) {
	defer wg.Done()
	var readString []string
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)

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
	storeString <- readString
}

func wordByWord(tword chan int, readStringChan chan []string) {
	defer wg.Done()

	total := 0
	readString := <-readStringChan
	re := regexp.MustCompile("[^\\s]+")
	for _, line := range readString {
		if len(line) != 0 {
			words := re.FindAllString(line, -1)
			total += len(words)
		}
	}
	tword <- total
}

func lineByLine(tline chan int, readStringChan chan []string) {
	defer wg.Done()
	readString := <-readStringChan
	totalL := len(readString)
	tline <- totalL
}

func charByChar(tchar chan int, readStringChan chan []string) {
	defer wg.Done()
	total := 0
	totalChars = 0
	readString := <-readStringChan
	for _, line := range readString {
		total += len(string(line))
	}
	tchar <- total
}

func printTotalResult(tline chan int, tword chan int, tchar chan int) {
	defer wg.Done()
	totalWords = <-tword
	totalChars = <-tchar
	totalLines = <-tline

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

func createFile(name string) {
	f, err := os.Create("/tmp/output.txt")
	if err != nil {
		fmt.Println("error occur in the creation of the file")
		f.Close()
		return
	}

	fmt.Println("File successfully created")
	f.Close()
}

func printToFile(tline chan int, tword chan int, tchar chan int) {
	defer wg.Done()
	totalLine := <-tline
	totalWord := <-tword
	totalChar := <-tchar
	f, err := os.OpenFile("/tmp/output.txt", os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("error occur")
		f.Close()
		return
	}
	fmt.Fprintf(f, "%d %d %d total\n", totalLine, totalWord, totalChar)
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: wcShaMem <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		stringChan := make(chan []string, 1000)
		wg.Add(1)
		go readFile(file, stringChan)
		wg.Wait()
		wg.Add(3)
		resultTotalLine := make(chan int, 5)
		resultTotalChar := make(chan int, 5)
		resultTotalWord := make(chan int, 5)
		go lineByLine(resultTotalLine, stringChan)
		go wordByWord(resultTotalWord, stringChan)
		go charByChar(resultTotalChar, stringChan)
		wg.Wait()
		wg.Add(1) // Add another goroutine to print the total result after all reading goroutines have finished.
		go printTotalResult(resultTotalLine, resultTotalWord, resultTotalChar)
		wg.Wait()
	}
}
