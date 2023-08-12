package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Enter a 4 bit binary number: ")

	userInput := captureInput(os.Stdin)

	fmt.Println("Your number is:", userInput)
}

func captureInput2() (input string) {
    // Don't bother with this; just experimenting w/ Scans
	_, err := fmt.Scanln(&input)
	if err != nil {
		if errors.Is(err, io.EOF)  {
		    fmt.Println("Reading file finished ...")
		}
	}
	return input
}

func captureInput(rdr io.Reader) (input string) {
	_, err := fmt.Fscanln(rdr, &input)
	if err != nil {
		panic(err)
	}

	return input
}


