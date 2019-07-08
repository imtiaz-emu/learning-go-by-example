package main

import (
	"fmt"
	"net/http"
)

func getContentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Length")
	if ctype == "" {
		return "", fmt.Errorf("Content-Length is not available")
	}
	return ctype, nil
}

func main() {
	url := "https://stackoverflow.com/questions/tagged/ruby-on-rails"
	result, err := getContentType(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
