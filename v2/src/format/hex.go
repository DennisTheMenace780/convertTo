package format

import (
	"convertTo/src/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

type Hex string

func (hs *Hex) ToDecimalString() Decimal {
	// Largest hex string that can be handled is
	// 7FFFFFFFFFFFFFFF for a 64 bit architecture.
	var sum float64
    strToProcess, _ := strings.CutPrefix(string(*hs), "0x")
	chunkedHexString := utils.Chunk(strToProcess, 1)

	exp := len(chunkedHexString) - 1

	for _, char := range chunkedHexString {
		if decNumAsStr, ok := HexToDecimalMap[char]; ok {
			sum += computeValue(decNumAsStr, exp)
		} else {
			sum += computeValue(char, exp)
		}
		exp -= 1
	}
	intSum := int(sum)
	return Decimal(strconv.Itoa(intSum))
}

func (hs *Hex) Display() {
    fmt.Println("0x" + *hs)
}

func IsHexFormat(str string) bool {
	prefix := str[0:2]
	return prefix == "0x"
}

func computeValue(str string, power int) float64 {
	numAsInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: ", err)
        os.Exit(1)
	}
	return float64(numAsInt) * math.Pow(16, float64(power))
}
