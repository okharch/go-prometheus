package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	// Define command-line flags
	baseURL := flag.String("url", "http://localhost:8080/allocate", "The URL of the backend allocate handler")
	size := flag.Int("size", 1048576, "Memory allocation size in bytes (default: 1 MB)")
	duration := flag.Int("duration", 5, "Duration of CPU simulation in seconds (default: 5 seconds)")

	// Parse command-line arguments
	flag.Parse()

	// Validate required parameters
	if *size <= 0 {
		fmt.Println("Error: size must be greater than 0")
		os.Exit(1)
	}
	if *duration <= 0 {
		fmt.Println("Error: duration must be greater than 0")
		os.Exit(1)
	}

	// Prepare request parameters
	params := url.Values{}
	params.Add("size", fmt.Sprintf("%d", *size))
	params.Add("duration", fmt.Sprintf("%d", *duration))

	// Make the POST request
	resp, err := http.PostForm(*baseURL, params)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Print response status
	fmt.Printf("Request completed successfully\n")
	fmt.Printf("Response Status: %s\n", resp.Status)
}
