// slices are like vectors in c++
// tey give us a lot of functionalities, like dynamic memory allocation and printing only a part of the vector

package main

import "fmt"

func main() {
	// slice initialization is similar to arrays just the fact that the size of the arrays isn't defined
	var fruitList = []string{} // here even an empty brackets (curly brackets need to be entered)
	// to add to the existing vector we use, append function
	fruitList = append(fruitList, "Apple", "Banana", "Grape", "Pineapple")
	// you can directly append an array at the end of an array using append
	// using
	// fruitList = append(fruitList, fruitList...)
	// fmt.Println(fruitList)

	// hwo to remove a value from a slice based on index
	// basically append the values before the given index followed by the values just after the specified index.

	var index int = 2
	fruitList = append(fruitList[:index], fruitList[index+1:]...)
	fmt.Println(fruitList)

}
