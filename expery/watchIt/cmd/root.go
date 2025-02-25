/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var (
	Directory   = os.Getenv("HOME")
	WatchedWord = "foo"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "watchIt",
	Short: "Specify a directory to watch",
	Long:  `Specify a directory to watch, or It will use your default HOME directory`,
	Run: func(cmd *cobra.Command, args []string) {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()
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
	rootCmd.PersistentFlags().StringVarP(&Directory, "directory", "d", Directory, "Use to specify the directory to watch otherwise it will used HOME directory")
	rootCmd.PersistentFlags().StringVarP(&WatchedWord, "watchedword", "w", WatchedWord, "Use to specify words to search")
}
