package main

import (
	"bytes"
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
		userInput, err := captureInput(os.Stdin)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Require a 4 digit binary number, please try again")
			}
		}

		_, ok := is4DigitBinaryString(userInput)

		if !ok {
			fmt.Println("Require a 4 digit binary number, please try again")
		} else {
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

func findPandQ(m float64) (float64, float64) {
	// solving 2^p < m
	p := math.Floor(math.Log2(m))
	return p, p + 1
}

func decimalToBinary(m string) string {

	var output bytes.Buffer

	v, err := strconv.Atoi(m)
	if err != nil {
		fmt.Println("Error", err)
	}

	topNumber := float64(v)
	p, _ := findPandQ(topNumber)

	for p > -1 {

		remainder := math.Mod(topNumber, math.Pow(2.0, p))

		if topNumber >= math.Pow(2.0, p) {
			output.WriteString("1")
		} else {
			output.WriteString("0")
		}

        topNumber = remainder

		p -= 1
	}
	return output.String()
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
