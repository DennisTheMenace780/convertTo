package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
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
		if bit != "1" && bit != "0" { // Equivalent !(A || B)
			validString = false
			break
		} else {
			validString = true
		}
	}
	return length, validString
}

func binaryToDecimal(binNum string) string {
	numStr := strings.Split(binNum, "")

	var decimalValue float64

	for i, v := range numStr {
		exponent := float64(len(numStr)) - (float64(i) + 1.0)
		n, err := strconv.Atoi(v)
		if err != nil {
            fmt.Println("Problem converting to string", err)
		}
		decimalValue += float64(n) * math.Pow(2.0, exponent)
	}

    return strconv.Itoa(int(decimalValue))
}

func captureInput2(f *os.File) (input string) {
	// Don't bother with this; just experimenting w/ Scans
	_, err := fmt.Scan(&input)
	if err != nil {
		if errors.Is(err, io.EOF) {
			fmt.Println("Reading file finished ...")
		}
	}
	return input
}
