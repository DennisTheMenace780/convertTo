package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func numLeadingZerosNeeded(binStr string) int {
	var numZeros int
    exp := 2.0
    lb, ub := 0, 4
	for {
		if len(binStr) > lb && len(binStr) < ub {
			numZeros = ub - len(binStr)
			break
		}
		lb = ub
		exp += 1
        ub = int(math.Pow(2.0, exp))
	}
    return numZeros
}

func NewBinaryString(binStr string) BinaryString {
	var b bytes.Buffer
    numZeros := numLeadingZerosNeeded(binStr)
	// fill in empty space with 0s
	for n := 0; n < numZeros; n++ {
		b.WriteString("0")
	}

	b.WriteString(binStr)
	return BinaryString{bstring: b.String()}
}

type BinaryString struct {
	bstring string
}

func (bs *BinaryString) chunkString(size int) []string {
	chunkSize := size
	if len(bs.bstring) == 0 {
		return nil
	}
	if chunkSize >= len(bs.bstring) {
		return []string{bs.bstring}
	}
	var chunks []string = make([]string, 0, (len(bs.bstring)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range bs.bstring {
		if currentLen == chunkSize {
			chunks = append(chunks, bs.bstring[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, bs.bstring[currentStart:])
	return chunks
}

func (bs *BinaryString) Display() {
	var b bytes.Buffer
	chunkedStr := bs.chunkString(8)
	for _, v := range chunkedStr {
		b.WriteString(v + " ")
	}
	s := strings.TrimRight(b.String(), " ")
	fmt.Println(s)
}

func (bs *BinaryString) ToDecimalString() DecimalString {
	str := strings.Split(bs.bstring, "")

	var decimalValue float64

	for i, v := range str {
		exponent := float64(len(str)) - (float64(i) + 1.0)
		n, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Problem converting to string", err)
		}
		decimalValue += float64(n) * math.Pow(2.0, exponent)
	}
	return DecimalString{dstring: strconv.Itoa(int(decimalValue))}
}

func isBinaryFormat(str string) bool {
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
