package main

import (
	"bufio"
	"fmt"
	"os"
)

// raeder variable
var reader = bufio.NewReader(os.Stdin)

// get user inputs from cli
func userInputs(text string) string {

	fmt.Print(text)
	userInput, _ := reader.ReadString('\n')
	return userInput
}
