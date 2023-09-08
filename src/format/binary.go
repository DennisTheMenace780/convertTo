package format

import (
	"binaryTo/src/utils"
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type BinaryString struct {
	Bstring string
}

func (bs *BinaryString) Display() {
	var bufr bytes.Buffer
	chunkedStr := utils.Chunk(bs.Bstring, 8)
	for _, str := range chunkedStr {
		bufr.WriteString(str + " ")
	}
	s := strings.TrimRight(bufr.String(), " ")
	fmt.Println(s)
}

func (bs *BinaryString) ToDecimalString() DecimalString {
	str := strings.Split(bs.Bstring, "")

	var decimalValue float64

	for i, v := range str {
		exponent := float64(len(str)) - (float64(i) + 1.0)
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Problem converting to string", err)
		}
		decimalValue += float64(n) * math.Pow(2.0, exponent)
	}
	return DecimalString{Dstring: strconv.Itoa(int(decimalValue))}
}

func NewBinaryString(binStr string) BinaryString {
	var bufr bytes.Buffer
	numZeros := numLeadingZerosNeeded(binStr)
	// fill in empty space with 0s
	for n := 0; n < numZeros; n++ {
		bufr.WriteString("0")
	}

	bufr.WriteString(binStr)
	return BinaryString{Bstring: bufr.String()}
}

func IsBinaryFormat(str string) bool {
	var ok bool
	// Maybe return error instead with a specific format string?
	splitStr := strings.Split(str, "")

	for _, bit := range splitStr {
		if bit != "1" && bit != "0" {
			ok = false
			break
		} else {
			ok = true
		}
	}
	return ok
}
func numLeadingZerosNeeded(binStr string) int {
	var numZeros int
	exponent := 2.0
	lbound, ubound := 0, 4
	for {
		if len(binStr) >= lbound && len(binStr) <= ubound {
			numZeros = ubound - len(binStr)
			break
		}
		lbound = ubound
		exponent += 1
		ubound = int(math.Pow(2.0, exponent))
	}
	return numZeros
}
