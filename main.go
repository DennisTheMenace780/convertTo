package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type CmdLineInputs struct {
	toBinaryShort  *bool
	toBinaryLong   *bool
	toDecimalShort *bool
	toDecimalLong  *bool
}

func NewCmdLineInputs() CmdLineInputs {
	inputs := CmdLineInputs{
		toBinaryShort:  flag.Bool("b", false, "convert a decimal number to binary"),
		toBinaryLong:   flag.Bool("binary", false, "convert a decimal number to binary"),
		toDecimalShort: flag.Bool("d", false, "convert a binary number to decimal"),
		toDecimalLong:  flag.Bool("decimal", false, "convert a binary number to decimal"),
	}

	inputs.parse()

	return inputs
}

func (c *CmdLineInputs) parse() {
	flag.Parse()
}

func (c *CmdLineInputs) validInput() error {

	if len(os.Args[1:]) != 2 {
		return errors.New("More than one flag or arg provided")
	}
	return nil
}

func (c *CmdLineInputs) convertToBinary() bool {
	return *c.toBinaryShort || *c.toBinaryLong
}

func (c *CmdLineInputs) convertToDecimal() bool {
	return *c.toDecimalShort || *c.toDecimalLong
}

func (c *CmdLineInputs) dispatchConversion(numToConvert string) {
	if c.convertToBinary() {
		fmt.Println("Converting", numToConvert, "to binary")
		prettyPrintBinaryString(decimalToBinary(numToConvert))
	}
	if c.convertToDecimal() {
		fmt.Println("Converting", numToConvert, "to decimal")
		if !isBinaryString(numToConvert) {
			panic("Invalid input")
		}
		fmt.Println(binaryToDecimal(numToConvert))
	}
}

func main() {
	inputs := NewCmdLineInputs()

	err := inputs.validInput()
	if err != nil {
		panic(err)
	}

	numToConvert := os.Args[2]
	inputs.dispatchConversion(numToConvert)
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

func constructBinaryString(binStr string) string {
	var b bytes.Buffer
	var diff int

	if len(binStr) < 4 {
		diff = 4 - len(binStr)
	} else if len(binStr) > 4 && len(binStr) < 8 {
		diff = 8 - len(binStr)
	} else if len(binStr) > 8 && len(binStr) < 16 {
		diff = 16 - len(binStr)
	} else if len(binStr) > 16 && len(binStr) < 32 {
        diff = 32 - len(binStr)
    }

    // fill in empty space with 0s
	for n := 0; n < diff; n++ {
		b.WriteString("0")
	}

	b.WriteString(binStr)

	return b.String()
}

func prettyPrintBinaryString(binStr string) {
	var b bytes.Buffer
	str := constructBinaryString(binStr)
	c := Chunks(str, 8)

	for _, v := range c {
		b.WriteString(v + " ")
	}
    s := strings.TrimRight(b.String(), " ")
    fmt.Println(s)
}

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}
