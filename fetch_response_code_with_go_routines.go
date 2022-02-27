package main

import (
	"fmt"
	"net/http"
	"sync"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	fmt.Println("-----")

	if err != nil {
		fmt.Printf("Error occured while fetching url : %v, error details : %v \n", url, err)
		return "", err
	}

	// Close the response body
	defer resp.Body.Close()

	contentType := resp.Header.Get("content-type")
	if contentType == "" {
		return "", fmt.Errorf("Can not find content-type header for url : %v \n", url)
	} else {
		fmt.Printf("Content Type : %v, for the URL: %v \n", contentType, url)
	}
	return contentType, nil
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"NoSuchUrl",
	}

	/*
	* Here, we have used WaitGroup from sync package, so we can ask main thread to wait until
	* all the go routines finishes their execution.
	 */
	var wg sync.WaitGroup
	for _, url := range urls {

		wg.Add(1)
		/**
		* Following line showcase the usage of go routine.
		* go routine will create separte thread for executing a function
		 */
		go func(url string) {
			contentType(url)
			wg.Done()
		}(url)
	}

	wg.Wait()
}
