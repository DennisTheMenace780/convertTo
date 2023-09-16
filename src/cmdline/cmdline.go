package cmdline

import (
	"errors"
	"flag"
	"fmt"
	"os"
    "convertTo/src/format"
)

type CmdLineInputs struct {
	toBinaryLong   *bool
	toBinaryShort  *bool
	toDecimalLong  *bool
	toDecimalShort *bool
	toHexFromDecimalLong  *bool
	toHexFromDecimalShort *bool
}

func NewCmdLineInputs() CmdLineInputs {
	inputs := CmdLineInputs{
		toBinaryLong:   flag.Bool("binary", false, ""),
		toBinaryShort:  flag.Bool("b", false, "convert a decimal number to binary"),
		toDecimalLong:  flag.Bool("decimal", false, ""),
		toDecimalShort: flag.Bool("d", false, " convert a binary number to decimal"),
		toHexFromDecimalLong:  flag.Bool("hex", false, ""),
		toHexFromDecimalShort: flag.Bool("h", false, " convert a decimal number to hexadecimal"),
	}
	flag.Parse()
	return inputs
}

func (c *CmdLineInputs) IsValid() error {
	if len(os.Args[1:]) != 2 {
		return errors.New("More than one flag or arg provided")
	}
    fmt.Println(os.Args)
	return nil
}

func (c *CmdLineInputs) convertToBinary() bool {
	return *c.toBinaryShort || *c.toBinaryLong
}

func (c *CmdLineInputs) convertToDecimal() bool {
	return *c.toDecimalShort || *c.toDecimalLong
}

func (c *CmdLineInputs) convertToHexadecimalFromDecimal() bool {
	return *c.toHexFromDecimalShort || *c.toHexFromDecimalLong
}

func (c *CmdLineInputs) DispatchConversion(numToConvert string) {

	if c.convertToBinary() {
		fmt.Println("Converting", numToConvert, "to binary")

		decimalStr := format.NewDecimalString(numToConvert)
		binStr := decimalStr.ToBinaryString()
		binStr.Display()
	}

	if c.convertToDecimal() {
		if !format.IsBinaryFormat(numToConvert) {
			fmt.Println("Invalid string")
			os.Exit(1)
		}
		fmt.Println("Converting", numToConvert, "to decimal")

		binStr := format.NewBinaryString(numToConvert)
		decimalStr := binStr.ToDecimalString()
		decimalStr.Display()
	}

    if c.convertToHexadecimalFromDecimal() {
        decimalStr := format.NewDecimalString(numToConvert)
        hexStr := decimalStr.ToHexString()
        hexStr.Display()
    }
}
