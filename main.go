package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

func main() {

	RunConsolePromptApp()

}

func captureInput(rdr io.Reader) (string, error) {
	var input string
	_, err := fmt.Fscanln(rdr, &input)
	if err != nil {
		return input, err
	}

	return input, err
}

func isBinaryString(str string) (ok bool) {
	// Maybe return error instead with a specific format string?
	splitStr := strings.Split(str, "")

	for _, bit := range splitStr {
		if bit != "1" && bit != "0" { // Equivalent !(A || B)
			ok = false
			break
		} else {
			ok = true
		}
	}
	return ok
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

func findLargestExponent(m float64) (p float64) {
	// solving for 2^p = m
	return math.Floor(math.Log2(m))
}

func decimalToBinary(m string) string {
	var binaryString bytes.Buffer

	v, err := strconv.Atoi(m)
	if err != nil {
		fmt.Println("Error", err)
	}

	numerator := float64(v)
	p := findLargestExponent(numerator)

	for p > -1 {
		denominator := math.Pow(2.0, p)
		remainder := math.Mod(numerator, denominator)

		if numerator >= denominator {
			binaryString.WriteString("1")
		} else {
			binaryString.WriteString("0")
		}

		numerator = remainder
		p -= 1
	}
	return binaryString.String()
}

func prettyBinaryString(binStr string) string {
	var b bytes.Buffer
	var diff int

	if len(binStr) < 4 {
		diff = 4 - len(binStr)
	} else if len(binStr) > 4 && len(binStr) < 8 {
		diff = 8 - len(binStr)
	}

	for n := 0; n < diff; n++ {
		b.WriteString("0")
	}
	b.WriteString(binStr)

	return b.String()
}
