package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

type DecimalString struct {
	dstring string
}

func NewDecimalString(decimalStr string) DecimalString {
	return DecimalString{dstring: decimalStr}
}

func (ds *DecimalString) ToBinaryString() BinaryString {
	var binStr bytes.Buffer
	val, err := strconv.Atoi(ds.dstring)
	if err != nil {
		fmt.Println("Error:", err)
	}

	numerator := float64(val)
	largestExponent := math.Floor(math.Log2(numerator))

	for largestExponent > -1 {
		denominator := math.Pow(2.0, largestExponent)
		remainder := math.Mod(numerator, denominator)

		if numerator >= denominator {
			binStr.WriteString("1")
		} else {
			binStr.WriteString("0")
		}

		numerator = remainder
		largestExponent -= 1
	}
    return BinaryString{bstring: binStr.String()}
}

func (ds *DecimalString) Display() {
    fmt.Println(ds.dstring)
}
