/*
This code groupes 3 exercises of chapter 7.  it will be done in step with commit and comments to illustarte the changes
  - use the functionality of byCharacter.go, byLine.go, and byWord.go in order to create a simplified version of the wc(1)
    UNIX utility.
  - Create a full version of the wc(1) UNIX utility, using the viper package to process command line options
  - Create a full version of the wc(1) UNIX utility, using commmands instead of command line options, with the
    help of the cobra package
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

/*
printResult takes two arguments
and then total only if the number of characters of the second arguments is zero
otherwise it prints both two arguments
*/
func countPerWordPerFile(s []string) (total int) {
	total = 0
	re := regexp.MustCompile("[^\\s]+")
	for _, line := range s {
		if len(line) != 0 {
			words := re.FindAllString(line, -1)
			total += len(words)
		}
	}
	return
}

func countPerCharacterPerFile(s []string) (total int) {
	total = 0
	for _, line := range s {
		total += len(string(line))
	}
	return
}

func readfile(fileName string) ([]string, error) {
	var linesReaded []string
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		// ReadString() returns two values: the string that was read and an error variable.
		line, err := r.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				linesReaded = append(linesReaded, line)
			}
			break
		}

		if err != nil {
			fmt.Printf("error reading file %s", err)
			return nil, err
		}
		linesReaded = append(linesReaded, line)
	}
	return linesReaded, nil
}

func main() {
	args := os.Args
	lengthArgs := len(args)
	var readLines []string
	pflag.BoolP("lines", "l", true, "--lines, -l to count lines")
	pflag.BoolP("chars", "c", true, "--chars, -c to count characters")
	pflag.BoolP("words", "w", true, "--words, -w to count words")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	linesActivated := viper.GetBool("lines")
	fmt.Println("viper is ", linesActivated)

	if lengthArgs > 1 {
		var totalLines int
		for index := range lengthArgs - 1 {
			readlinesPerfile, _ := readfile(args[index+1])
			totalLines += len(readlinesPerfile)
			fmt.Printf("Total number of lines: %d %s\n", len(readlinesPerfile), args[index+1])
		}
		if lengthArgs > 2 {
			fmt.Println("Total number of lines:", totalLines)
		}

		var totalWords int
		for index := range lengthArgs - 1 {
			readlinesPerfile, _ := readfile(args[index+1])
			totalwordsPerfile := countPerWordPerFile(readlinesPerfile)
			fmt.Printf("Total number of Words: %d %s\n", totalwordsPerfile, args[index+1])
			totalWords += totalwordsPerfile
		}
		if lengthArgs > 2 {
			fmt.Println("Total number of Words:", totalWords)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			readLines = append(readLines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		fmt.Printf("Total number of lines: %d\n", len(readLines))
		totalWords := countPerWordPerFile(readLines)
		fmt.Printf("Total number of Words: %d\n", totalWords)
		totalChars := countPerCharacterPerFile(readLines)
		fmt.Printf("Total number of Characters: %d\n", totalChars)

	}
}
