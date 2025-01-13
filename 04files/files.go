package main

import (
	"fmt"
	"io"
	"os"
)

// working with files requires a io utils package and defer once the work with files is done

func main() {
	content := "This is some new content to be added to files."
	file, err := os.Create("./newfile.txt")
	if err != nil {
		panic(err)
	}
	// for the first time initilization use the walrus operator but from the next time simply use the assignment operator.
	_, err = io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("The content of the file", readFile("./newfile.txt"))
}

//  creation is a OS operation so we'll use methods under the OS packaeg but the reading and other operations are coming under the utility of ioutil.
// new update, ioutil is depreceated.
// files by default are not read in a string format

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
