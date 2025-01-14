package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// create some body payload, (using NewReader)
	// define a url for posting to
	// once the request is sent, catch the response and make appropriate response.
	const myurl = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
		{
			"coursename":"ReactJS",
			"somerandomKey":"1"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody) //url to send to, the content type, and the post body payload
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
