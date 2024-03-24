package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 || os.Args[0] != "banks" || os.Args[1] != "is" || os.Args[2] != "gay" {
		fmt.Println("Usage: banks is gay")
		return
	}

	fmt.Println("Yes, he is!")
}
