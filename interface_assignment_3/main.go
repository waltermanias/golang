package main

import (
	"fmt"
	"os"
)

func main() {

	filename := os.Args[1]

	sb, error := os.ReadFile(filename)
	if error != nil {
		os.Exit(1)
	}

	fmt.Println(string(sb))

}
