package main

import "fmt"

// there is a concept of immediately involked functions, too or anonymous functions;
// you cannot define functions inside functions

func main() {
	store := proAddr(2, 3, 4, 55, 6, 7, 8, 8)
	fmt.Println("Sum is ", store)
}

// we can specify ... to make a function have any number of inputs but their tyoe should be defined correctly.
// Note, a function can return more than one values too.

func proAddr(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
