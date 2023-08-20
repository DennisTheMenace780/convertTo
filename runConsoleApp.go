package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func RunConsolePromptApp() {

	var userInput string

	for {
		input, err := captureInput(os.Stdin)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Require a 4 digit binary number, please try again")
			}
		}

		ok := isBinaryString(input)

		if !ok {
			fmt.Println("Require a 4 digit binary number, please try again")
		} else {
			userInput = input
			break
		}
	}

	fmt.Println("Your number is:", userInput)
}
