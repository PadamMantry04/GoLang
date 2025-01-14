package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string   `json:"coursename"` // note here coursename is the alias that shall be visible when this data is seen.
	Price    int      `json:"paisa"`
	Password string   `json:"-"` // if you state a hyphen, then the field gets omitted or not shown
	Tags     []string `json:"omitempty"`
}

func main() {
	lcoCourses := []courses{
		{"ReactJS", 500, "pm123", []string{"Tag1", "Tag2"}},
		{"NodeJS", 500, "pm122", []string{"Tag1", "Tag2"}},
		{"NestJs", 500, "pm14423", []string{"Tag1", "Tag2"}},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // use Marshall indent to indent or pretty print the json data
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finalJson)
}
