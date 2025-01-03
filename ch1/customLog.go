package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	fmt.Println(LOGFILE)
	// The call to os.OpenFile() creates the log file for writing,
	// if it does not already exist, or opens it for writing
	// by appending new data at the end of it (os.O_APPEND)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// The defer keyword tells Go to execute the statement just before the current function returns.
	defer f.Close()
	iLog := log.New(f, "ilog", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go 4th edition!")

}
