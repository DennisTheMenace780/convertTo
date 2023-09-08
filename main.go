package main

import (
	"binaryTo/src/cmdline"
	"os"
)

func main() {
	inputs := cmdline.NewCmdLineInputs()

	err := inputs.IsValid()
	if err != nil {
		panic(err)
	}

	numToConvert := os.Args[2]
	inputs.DispatchConversion(numToConvert)
}

