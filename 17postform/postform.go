package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// create some body payload, (using NewReader)
	// define a url for posting to
	// once the request is sent, catch the response and make appropriate response.
	const myurl = "http://localhost:8000/post"
	data := url.Values{}
	data.Add("firstname", "padam")
	data.Add("lastname", "mantry")
	data.Add("email", "padam@gmail.com")

	response, err := http.PostForm(myurl, data) //url to send to and the postform body payload
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
