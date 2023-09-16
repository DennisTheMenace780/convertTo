package format

import (
	"bytes"
	"convertTo/v2/src/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Binary string

func (bs *Binary) Display() {
    // Break the functionality up in here, the Display is doing too much.
	var bufr bytes.Buffer
	s := string(*bs)

	// fill in empty space with 0s
	numZeros := prefixZeros(s)
	for n := 0; n < numZeros; n++ {
		bufr.WriteString("0")
	}
    bufr.WriteString(s)

	chunkedStr := utils.Chunk(bufr.String(), 8)
    bufr.Reset()
	for _, str := range chunkedStr {
		bufr.WriteString(str + " ")
	}
	trimStr := strings.TrimRight(bufr.String(), " ")
	fmt.Println(trimStr)
}

func (bs *Binary) ToDecimalString() Decimal {
	splitStr := strings.Split(string(*bs), "")
	var sum float64

	for i, v := range splitStr {
		exp := float64(len(splitStr)) - (float64(i) + 1.0)
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Problem converting to string", err)
		}
		sum += float64(n) * math.Pow(2.0, exp)
	}
	return Decimal(strconv.Itoa(int(sum)))
}

func (bs *Binary) ToString() string {
    return string(*bs)
}

func NewBinaryString(binStr string) Binary {
	var bufr bytes.Buffer
	numZeros := prefixZeros(binStr)
	// fill in empty space with 0s
	for n := 0; n < numZeros; n++ {
		bufr.WriteString("0")
	}

	bufr.WriteString(binStr)

	return Binary(bufr.String())
}

func IsBinaryFormat(str string) bool {
	var ok bool
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

func prefixZeros(binStr string) int {
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
