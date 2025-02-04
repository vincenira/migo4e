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

func printTotalResult(results ...int) {
	if results[0] != 0 {
		fmt.Printf("%d  ", results[0])
	}
	if results[1] != 0 {
		fmt.Printf("%d ", results[1])
	}
	if results[2] != 0 {
		fmt.Printf("%d ", results[2])
	}
	fmt.Printf("total\n")
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
	var readLines []string
	pflag.BoolP("lines", "l", false, "--lines, -l to count lines")
	pflag.BoolP("chars", "c", false, "--chars, -c to count characters")
	pflag.BoolP("words", "w", false, "--words, -w to count words")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	linesActivated := viper.GetBool("lines")
	wordsActivated := viper.GetBool("words")
	charsActivated := viper.GetBool("chars")
	args := pflag.Args()
	lengthArgs := len(args)
	if lengthArgs > 0 {
		var totalLines int
		var totalWords int
		var totalChars int
		for index := range lengthArgs {
			readlinesPerfile, _ := readfile(args[index])
			linePerFile := len(readlinesPerfile)
			if linesActivated {
				totalLines += linePerFile
				fmt.Printf("%d  ", linePerFile)
			}
			if wordsActivated {
				totalwordsPerfile := countPerWordPerFile(readlinesPerfile)
				totalWords += totalwordsPerfile
				fmt.Printf("%d ", totalwordsPerfile)
			}

			if charsActivated {
				totalCharsPerfile := countPerCharacterPerFile(readlinesPerfile)
				totalChars += totalCharsPerfile
				fmt.Printf("%d ", totalCharsPerfile)
			}
			fmt.Printf("%s\n", args[index])
		}
		if lengthArgs > 1 {
			printTotalResult(totalLines, totalWords, totalChars)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			readLines = append(readLines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

		if linesActivated {
			fmt.Printf("%d  ", len(readLines))
		}
		if wordsActivated {
			totalWords := countPerWordPerFile(readLines)
			fmt.Printf("%d ", totalWords)
		}
		if charsActivated {
			totalChars := countPerCharacterPerFile(readLines)
			fmt.Printf("%d ", totalChars)
		}
		fmt.Printf("\n")
	}
}
