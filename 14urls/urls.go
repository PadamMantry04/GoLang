package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://example.com/api/resource?param1=value1&param2=value2&param3=value3&param4=value4&param5=value5&param6=value6&param7=value7&param8=value8&param9=value9&param10=value10"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	// fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	qparams := result.Query() // this makes the different queries as key value pairs in the url provided.
	fmt.Println(qparams)

	for k, v := range qparams {
		fmt.Println(k, v)
	}

	// note you can modify the existing url using or create a new url using
	partsofURL := &url.URL{
		Scheme: "https",
		Host:   "york.com",
		Path:   "/api/v1",
	}

	newurl1 := partsofURL.String()
	fmt.Println(newurl1)
}
