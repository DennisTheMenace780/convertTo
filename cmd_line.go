package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type CmdLineInputs struct {
	toBinaryLong   *bool
	toBinaryShort  *bool
	toDecimalLong  *bool
	toDecimalShort *bool
}

func NewCmdLineInputs() CmdLineInputs {
	inputs := CmdLineInputs{
		toBinaryLong:   flag.Bool("binary", false, ""),
		toBinaryShort:  flag.Bool("b", false, "convert a decimal number to binary"),
		toDecimalLong:  flag.Bool("decimal", false, ""),
		toDecimalShort: flag.Bool("d", false, " convert a binary number to decimal"),
	}
	flag.Parse()
	return inputs
}

func (c *CmdLineInputs) isValid() error {
	if len(os.Args[1:]) != 2 {
		return errors.New("More than one flag or arg provided")
	}
	return nil
}

func (c *CmdLineInputs) convertToBinary() bool {
	return *c.toBinaryShort || *c.toBinaryLong
}

func (c *CmdLineInputs) convertToDecimal() bool {
	return *c.toDecimalShort || *c.toDecimalLong
}

func (c *CmdLineInputs) dispatchConversion(numToConvert string) {

	if c.convertToBinary() {
		fmt.Println("Converting", numToConvert, "to binary")

        decimalStr := NewDecimalString(numToConvert)
        binStr := decimalStr.ToBinaryString()
        binStr.Display()
	}

	if c.convertToDecimal() {
		if !isBinaryFormat(numToConvert) {
			fmt.Println("Invalid string")
			os.Exit(1)
		}
		fmt.Println("Converting", numToConvert, "to decimal")

		binStr := NewBinaryString(numToConvert)
        decimalStr := binStr.ToDecimalString()
        decimalStr.Display()
	}
}
