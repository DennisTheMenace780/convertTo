package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a 4 bit binary number: ")

	var userInput string

	for {
		userInput = captureInput(os.Stdin)

		_, ok := is4DigitBinaryString(userInput)

		if !ok {
			fmt.Println("Require a 4 digit binary number, please try again")
		} else {
			break
		}
	}

	fmt.Println("Your number is:", userInput)

}

func captureInput(rdr io.Reader) (input string) {
	_, err := fmt.Fscanln(rdr, &input)
	if err != nil {
        // TODO: #001 not sure how to handle "unexpected newline" error
		fmt.Println("Error:", err)
	}

	return input
}

func is4DigitBinaryString(input string) (length int, validString bool) {
	splitInput := strings.Split(input, "")
	length = len(splitInput)

	if length != 4 {
		validString = false
		return length, validString
	}

	for _, bit := range splitInput {
		if bit == "1" || bit == "0" { // Equivalent !(A || B)
			validString = true
		} else {
            break
        }
	}
	return length, validString
}

func captureInput2() (input string) {
	// Don't bother with this; just experimenting w/ Scans
	_, err := fmt.Scanln(&input)
	if err != nil {
		if errors.Is(err, io.EOF) {
			fmt.Println("Reading file finished ...")
		}
	}
	return input
}
