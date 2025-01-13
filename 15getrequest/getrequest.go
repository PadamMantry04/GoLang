package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Define the URL
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // Ensure the response body is closed

	// Check the status code
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", response.StatusCode)
		return
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// body1 := string(body)
	// Print the response body
	// for k, v := range body1 {
	// 	fmt.Println(k, v)
	// }
	fmt.Println(string(body))
}
