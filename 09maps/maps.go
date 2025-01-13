package main

import "fmt"

func main() {
	// lets discuss about the various functionalities of maps
	// 1. how to initialize a map
	// 2. how to add elements to a
	// 3. how to delete elements from a map
	// 4. how to loop through a map

	m1 := make(map[string]string)
	m1["1"] = "one"
	m1["2"] = "two"
	m1["3"] = "three"
	m1["4"] = "four"
	delete(m1, "1")
	for k, v := range m1 {
		fmt.Printf("key is %v and value is %v\n", k, v)
	}

}
