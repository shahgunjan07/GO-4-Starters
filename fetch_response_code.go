package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	fmt.Println("-----")

	if err!= nil {
		fmt.Printf("Error occured while fetching url : %v, error details : %v \n", url, err)	
		return  "", err
	} 

	// Close the response body
	defer resp.Body.Close()

	contentType := resp.Header.Get("content-type")
	if contentType == "" {
		return "", fmt.Errorf("Can not find content-type header for url : %v \n", url)
	}
	return contentType, nil
}

func main() {
	url := "https://www.google.com/"
	resp, err := contentType(url);

	if err != nil {
		fmt.Printf("Error generated while trying to connect :%v \n", url)
	} else {
		fmt.Printf("Content-Type received :%v \n ", resp)
	}

	url2 := "https://noSuchUrlExist.com/"
	resp2, err := contentType(url2);

	if err != nil {
		fmt.Printf("Error generated while trying to connect :%v \n", url2)
	} else {
		fmt.Printf("Content-Type received :%v \n ", resp2)
	}
}