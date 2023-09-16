package format

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Decimal string

func (ds *Decimal) ToBinaryString() Binary {
	var bufr bytes.Buffer
	val, err := strconv.Atoi(string(*ds))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	numerator := float64(val)
	exp := math.Floor(math.Log2(numerator))

	if exp == math.Inf(-1) { // Handle log(0)
		bufr.WriteString("0")
	}
	for exp > -1 {
		// Could also write this recursively if
		denominator := math.Pow(2.0, exp)
		remainder := math.Mod(numerator, denominator)

		if numerator >= denominator {
			bufr.WriteString("1")
		} else {
			bufr.WriteString("0")
		}

		numerator = remainder
		exp -= 1
	}
	return Binary(bufr.String())
}

func (ds *Decimal) ToHexString() Hex {
	var bufr bytes.Buffer
	val, err := strconv.Atoi(string(*ds))
	if err != nil {
		fmt.Println("Error:", err)
	}
	numerator := float64(val)
	var quotient float64
	// Found by solving 16^m = n (m < n, Floor(m) >= 0)
	largestExponent := math.Floor(0.25 * math.Log2(numerator))

	if largestExponent == math.Inf(-1) {
		bufr.WriteString("0")
        return Hex(bufr.String())
	}

	for largestExponent > -1 {
		denominator := math.Pow(16.0, largestExponent)
		remainder := math.Mod(numerator, denominator)
		quotient = math.Floor(numerator / denominator)

		strQuotient := strconv.Itoa(int(quotient))
		if hexNumAsStr, ok := DecimalToHexMap[strQuotient]; ok {
			bufr.WriteString(hexNumAsStr)
		} else {
			bufr.WriteString(strQuotient)
		}

		numerator = remainder
		largestExponent -= 1
	}
	return Hex(bufr.String())
}

func IsDecimalFormat(str string) bool {
	var ok bool
	_, err := strconv.Atoi(str)
	if err != nil {
		ok = false
		return ok
	}
	ok = true
	return ok
}

func (ds *Decimal) Display() {
	fmt.Println(*ds)
}
