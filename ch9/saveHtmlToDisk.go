// Modify wwwClient.go to save the HTML output to an external file.
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}
	/*
	   The url.Parse() function parses a string into a URL structure. This means that if the given ar-
	   gument is not a valid URL, url.Parse() is going to notice.
	*/
	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing:", err)
		return
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	/*
	   The http.NewRequest() function returns an http.Request object when provided with a method,
	   a URL, and an optional body. The http.MethodGet parameter defines that we want to retrieve
	   the data using a GET HTTP method, whereas URL.String() returns the string value of an http.
	   URL variable.
	*/
	request, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		fmt.Println("Get:", err)
		return
	}
	/*
	   The http.Do() function sends an HTTP request (http.Request) using an http.Client and gets
	   an http.Response back. So http.Do() does the job of http.Get() in a more detailed way
	*/
	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println("Error in Do():", err)
		return
	}
	/*
	   httpData.Status holds the HTTP status code of the response—this is important because it allows
	   you to understand what really happened with the request
	*/
	fmt.Println("Status code:", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Print(string(header))
	/*
	   The httputil.DumpResponse() function is used here to get the response from the server and
	   is mainly used for debugging purposes. The second argument of httputil.DumpResponse()
	   is a Boolean value that specifies whether the function is going to include the body or not in its
	   output—in our case, it is set to false, which excludes the response body from the output and
	   only prints the header. If you want to do the same on the server side, you should use httputil.
	   DumpRequest()
	*/
	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set:", characterSet[1])
	}

	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	/*
	   In the last part of the program, we use a technique for discovering the size of the server HTTP
	   response on our own. If we wanted to display the HTML output on our screen, we could have
	   printed the contents of the r buffer variable
	*/
	length := 0
	var buffer [1024]byte
	response, _ := io.ReadAll(httpData.Body)
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Calculated response data length:", length)
	fileName := "/tmp/saveHtml.html"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error occurred while creating the file")
		f.Close()
		return
	}
	f, err = os.OpenFile("/tmp/saveHtml.html", os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("error occur")
		f.Close()
		return
	}
	// fmt.Fprintf(f, "%s", response)
	fmt.Println(response)
	f.Close()
}
