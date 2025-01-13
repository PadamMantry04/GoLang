// there aren't much options for arrays in golang
// you have to initalize the space reserved for the array at the time of initialization and also specify the type.

package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	// consider an example of initialization followed by assignment.
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	// fruitList[2] = "Orange"
	fruitList[3] = "Grapes"

	var vegList = [4]string{"apple"}
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	fmt.Println(vegList)
	fmt.Println(len(vegList))
}
