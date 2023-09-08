package format

import (
	"binaryTo/src/utils"
	"fmt"
	"math"
	"strconv"
)

var DecimalToHexMap = map[string]string{
	"10": "A",
	"11": "B",
	"12": "C",
	"13": "D",
	"14": "E",
	"15": "F",
}

var HexToDecimalMap = map[string]string{
	"A": "10",
	"B": "11",
	"C": "12",
	"D": "13",
	"E": "14",
	"F": "15",
}

func NewHexString(hexStr string) HexString {
	return HexString{hstring: hexStr}
}

type HexString struct {
	hstring string
}

func (hs *HexString) Display() {
	fmt.Println("0x" + hs.hstring)
}

func (hs *HexString) ToDecimalString() DecimalString {
	// Largest hex string that can be handled is
	// 7FFFFFFFFFFFFFFF for a 64 bit architecture.
	var sum float64
	chunkedHexString := utils.Chunk(hs.hstring, 1)
	largestExponent := len(chunkedHexString) - 1

	for _, char := range chunkedHexString {
		if decNumAsStr, ok := HexToDecimalMap[char]; ok {
			sum += computeValue(decNumAsStr, largestExponent)
		} else {
			sum += computeValue(char, largestExponent)
		}
		largestExponent -= 1
	}
	return DecimalString{Dstring: strconv.Itoa(int(sum))}
}

func computeValue(str string, power int) float64 {
	numAsInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return float64(numAsInt) * math.Pow(16, float64(power))
}
