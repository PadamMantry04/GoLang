package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", text)
}
