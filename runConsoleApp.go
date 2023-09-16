package main

import (
	"convertTo/src/format"
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

		ok := format.IsBinaryFormat(input)

		if !ok {
			fmt.Println("Require a 4 digit binary number, please try again")
		} else {
			userInput = input
			break
		}
	}

	fmt.Println("Your number is:", userInput)
}

func captureInput(rdr io.Reader) (string, error) {
	var input string
	_, err := fmt.Fscanln(rdr, &input)
	if err != nil {
		return input, err
	}

	return input, err
}
