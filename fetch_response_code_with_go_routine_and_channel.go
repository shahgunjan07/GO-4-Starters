package main

import (
	"fmt"
	"net/http"
)

func contentType(url string, out chan string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		out <- "Error occured while fetching url : " + url + " , error details : " + err.Error()
		return "", err
	}

	// Close the response body
	defer resp.Body.Close()

	contentType := resp.Header.Get("content-type")
	if contentType == "" {
		out <- "Can not find content-type header for url : " + url
		return "", fmt.Errorf("Can not find content-type header for url : %v \n", url)
	} else {
		out <- "Content Type :" + contentType + ", for the URL:" + url
	}
	return contentType, nil
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"NoSuchUrl",
	}

	// Create a channel to hold output details
	c := make(chan string)
	for _, url := range urls {

		/**
		* Following line showcase the usage of go routine.
		* go routine will create separte thread for executing a function
		 */
		go func(url string) {
			contentType(url, c)
		}(url)
	}

	for range urls {
		out := <-c
		fmt.Println(out)
	}
}
