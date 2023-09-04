package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)


func main() {
	inputs := NewCmdLineInputs()

	err := inputs.isValid()
	if err != nil {
		panic(err)
	}

	numToConvert := os.Args[2]
	inputs.dispatchConversion(numToConvert)
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

