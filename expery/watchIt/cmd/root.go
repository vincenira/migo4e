/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"strings"

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
		watchDog, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watchDog.Close()
		done := make(chan bool)
		go func() {
			for {
				select {
				case event, ok := <-watchDog.Events:
					if !ok {
						return
					}
					if strings.Contains(strings.ToLower(WatchedWord), strings.ToLower(event.Name)) && event.Op&fsnotify.Create == fsnotify.Create || fsnotify.Write == fsnotify.Write || fsnotify.Remove == fsnotify.Remove || fsnotify.Rename == fsnotify.Rename {
						log.Printf("%s event was triggered on file: %s", event, event.Name)
					}
				case errWatchDog, ok := <-watchDog.Errors:
					if !ok {
						return
					}
					log.Printf("error: %v", errWatchDog)
				}
			}
		}()
		err = watchDog.Add(Directory)
		if err != nil {
			log.Fatal(err)
		}
		<-done // Wait for watcher to be closed.
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
