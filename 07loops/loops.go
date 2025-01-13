package main

import "fmt"

func main() {
	// we'll discuss the differernt concepts of loops in this regard.
	// the generic syntax is the same using k,v for maps or i,v for arrays or slices.
	// Apart from that the syntax is about continue and break and goto.
	i := 1
	for i < 10 {
		if i == 3 {
			goto lco
		}
		fmt.Println(i)
		i++ // here we primarily only use postfix increment operator.
	}

lco:
	//this is a goto label.
	fmt.Println("Just printing out the statements")
}
