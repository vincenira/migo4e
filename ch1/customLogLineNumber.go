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
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	/*
	   The desired functionality is implemented with the use of log.Lshortfile in the parameters of
	   log.New() or SetFlags(). The log.Lshortfile flag adds the filename as well as the line number
	   of the Go statement that printed the log entry in the log entry itself. If you use log.Llongfile
	   instead of log.Lshortfile, then you get the full path of the Go source file—usually, this is not
	   necessary, especially when you have a really long path.
	*/
	LstdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum ", LstdFlags)
	iLog.Println("Mastering Go, 4th edition!")
	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Printf("Another log entry")
}
