package main

import (
	"fmt"
	"os" // This package is used to interact with the operating system, including reading terminal input.
)

func main(){
	// os.Args is a built-in Go variable that contains a list of inputs from the terminal.
    // Let's check: is the number of inputs less than 2? if so then throw an error
    // (Index 0 = program name, Index 1 = username)
	if len(os.Args) < 2 {
		fmt.Println("Error: You must enter your username!")
		fmt.Println("Example: go run main.go 'username'")
	}

	// if the input is correct, than take the second input (index 1)
	username := os.Args[1]

	// Print the result for testing
	fmt.Printf("Searching github activity for user: %s/n",username)
}