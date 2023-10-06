package net

import (
	"net/http"
	"time"
)

func IsOnline() bool {
	// Create a timeout for the request
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// Make a GET request to a reliable server
	resp, err := client.Get("https://www.example.com")
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return true
	}
	return false
}
