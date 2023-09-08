package format

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

type DecimalString struct {
	Dstring string
}

func NewDecimalString(decStr string) DecimalString {
	return DecimalString{Dstring: decStr}
}

func (ds *DecimalString) ToBinaryString() BinaryString {
	var binStr bytes.Buffer
	val, err := strconv.Atoi(ds.Dstring)
	if err != nil {
		fmt.Println("Error:", err)
	}

	numerator := float64(val)
	largestExponent := math.Floor(math.Log2(numerator))

	if largestExponent == math.Inf(-1) {
		binStr.WriteString("0")
	}

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
	return BinaryString{Bstring: binStr.String()}
}

func (ds *DecimalString) ToHexString() HexString {
	var hexBufr bytes.Buffer
	val, err := strconv.Atoi(ds.Dstring)
	if err != nil {
		fmt.Println("Error:", err)
	}

	numerator := float64(val)
	var quotient float64
	// Found by solving 16^m = n (m < n, Floor(m) >= 0) 
	largestExponent := math.Floor(0.25 * math.Log2(numerator))

	if largestExponent == math.Inf(-1) {
		hexBufr.WriteString("0")
	}
	for largestExponent > -1 {
		denominator := math.Pow(16.0, largestExponent)
		remainder := math.Mod(numerator, denominator)
		quotient = numerator / denominator

		strQuotient := strconv.Itoa(int(quotient))
		if hexNumAsStr, ok := DecimalToHexMap[strQuotient]; ok {
			hexBufr.WriteString(hexNumAsStr)
		} else {
			hexBufr.WriteString(strQuotient)
		}

		numerator = remainder
		largestExponent -= 1
	}
	return HexString{hstring: hexBufr.String()}
}

func (ds *DecimalString) Display() {
	fmt.Println(ds.Dstring)
}
