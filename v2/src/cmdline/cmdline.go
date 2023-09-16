package cmdline

import (
	"convertTo/v2/src/format"
	"errors"
	"flag"
	"fmt"
	"os"
)

type CmdLineInputs struct {
	toBinaryLong          *bool
	toBinaryShort         *bool
	toDecimalLong         *bool
	toDecimalShort        *bool
	toHexFromDecimalLong  *bool
	toHexFromDecimalShort *bool
}

func NewCmdLineInputs() CmdLineInputs {
	inputs := CmdLineInputs{
		toBinaryLong:          flag.Bool("binary", false, ""),
		toBinaryShort:         flag.Bool("B", false, "convert a decimal number to binary"),
		toDecimalLong:         flag.Bool("decimal", false, ""),
		toDecimalShort:        flag.Bool("D", false, " convert a binary number to decimal"),
		toHexFromDecimalLong:  flag.Bool("hexadecimal", false, ""),
		toHexFromDecimalShort: flag.Bool("H", false, " convert a decimal number to hexadecimal"),
	}
	flag.Parse()
	return inputs
}

func (c *CmdLineInputs) IsValid() error {
	if len(os.Args[1:]) != 2 {
		return errors.New("More than one flag or arg provided")
	}
	return nil
}

func (c *CmdLineInputs) Dispatch(toConvert string) {

	switch {
	// Need to check if we're getting a Hex or a Decimal.
	case *c.toBinaryShort || *c.toBinaryLong:

		fmt.Println("Converting", toConvert, "to binary format")

		if format.IsHexFormat(toConvert) {
			fmt.Println("Error: Haven't handled HexFormatting Yet")
			os.Exit(1)
		}
		decimalStr := format.Decimal(toConvert)
		binaryStr := decimalStr.ToBinaryString()
		binaryStr.Display()

	case *c.toDecimalShort || *c.toDecimalLong:
		fmt.Println("Converting", toConvert, "to decimal")

		if format.IsBinaryFormat(toConvert) || format.IsDecimalFormat(toConvert) {
			// If the provided format fits the Binary format, then assume
			// the user intended to provide a decimal number (e.g., 1010 as decimal)
			decimalStr := format.Decimal(toConvert)
			decimalStr.Display()
			os.Exit(1)
		} else if format.IsHexFormat(toConvert) {
			hexStr := format.Hex(toConvert)
			decStr := hexStr.ToDecimalString()
			decStr.Display()
			os.Exit(1)
		} else {
			fmt.Println("invalid syntax")
		}

	case *c.toHexFromDecimalShort || *c.toHexFromDecimalLong:
		fmt.Println("Converting", toConvert, "to Hex format")

		decimalStr := format.Decimal(toConvert)
		hexStr := decimalStr.ToHexString()
		hexStr.Display()
	}
}
