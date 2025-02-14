/*
Copyright © 2025 vincenira <vincenira@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	wcBufChan "wcBufChan/helper"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wcBufChan",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var readLines []string
		linesActivated, _ := cmd.Flags().GetBool("lines")
		wordsActivated, _ := cmd.Flags().GetBool("words")
		charsActivated, _ := cmd.Flags().GetBool("chars")
		lengthArgs := len(args)
		if lengthArgs > 0 {
			var totalLines int
			var totalWords int
			var totalChars int
			for index := range lengthArgs {
				readlinesPerfile, _ := wcBufChan.Readfile(args[index])
				linePerFile := len(readlinesPerfile)
				if linesActivated {
					totalLines += linePerFile
					fmt.Printf("%d  ", linePerFile)
				}
				if wordsActivated {
					resultTotalChan := make(chan int, 5)
					go wcBufChan.CountPerWordPerFile(readlinesPerfile, resultTotalChan)
					totalwordsPerfile := <-resultTotalChan
					totalWords += totalwordsPerfile
					fmt.Printf("%d ", totalwordsPerfile)
				}

				if charsActivated {
					resultTotalCharsChan := make(chan int, 5)
					go wcBufChan.CountPerCharacterPerFile(readlinesPerfile, resultTotalCharsChan)
					totalCharsPerfile := <-resultTotalCharsChan
					totalChars += totalCharsPerfile
					fmt.Printf("%d ", totalCharsPerfile)
				}
				fmt.Printf("%s\n", args[index])
			}
			if lengthArgs > 1 {
				wcBufChan.PrintTotalResult(totalLines, totalWords, totalChars)
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
				resultTotalWordsChan := make(chan int, 5)
				go wcBufChan.CountPerWordPerFile(readLines, resultTotalWordsChan)
				totalWords := <-resultTotalWordsChan
				fmt.Printf("%d ", totalWords)
			}
			if charsActivated {
				resultTotalCharsChan := make(chan int, 5)
				go wcBufChan.CountPerCharacterPerFile(readLines, resultTotalCharsChan)
				totalChars := <-resultTotalCharsChan
				fmt.Printf("%d ", totalChars)
			}
			fmt.Printf("\n")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wcBufChan.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("lines", "l", false, "Counts the number of lines")
	rootCmd.Flags().BoolP("words", "w", false, "Counts the number of words")
	rootCmd.Flags().BoolP("chars", "c", false, "Counts the number of characters")
}
