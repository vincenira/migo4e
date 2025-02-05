/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"wcCobra/wcRewrite"

	"github.com/spf13/cobra"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
				readlinesPerfile, _ := wcRewrite.Readfile(args[index])
				linePerFile := len(readlinesPerfile)
				if linesActivated {
					totalLines += linePerFile
					fmt.Printf("%d  ", linePerFile)
				}
				if wordsActivated {
					totalwordsPerfile := wcRewrite.CountPerWordPerFile(readlinesPerfile)
					totalWords += totalwordsPerfile
					fmt.Printf("%d ", totalwordsPerfile)
				}

				if charsActivated {
					totalCharsPerfile := wcRewrite.CountPerCharacterPerFile(readlinesPerfile)
					totalChars += totalCharsPerfile
					fmt.Printf("%d ", totalCharsPerfile)
				}
				fmt.Printf("%s\n", args[index])
			}
			if lengthArgs > 1 {
				wcRewrite.PrintTotalResult(totalLines, totalWords, totalChars)
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
				totalWords := wcRewrite.CountPerWordPerFile(readLines)
				fmt.Printf("%d ", totalWords)
			}
			if charsActivated {
				totalChars := wcRewrite.CountPerCharacterPerFile(readLines)
				fmt.Printf("%d ", totalChars)
			}
			fmt.Printf("\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(countCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	countCmd.Flags().BoolP("lines", "l", false, "Counts the number of lines")
	countCmd.Flags().BoolP("words", "w", false, "Counts the number of words")
	countCmd.Flags().BoolP("chars", "c", false, "Counts the number of characters")
}
