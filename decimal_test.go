package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestDecimalString(t *testing.T) {

	t.Run("Creating a NewBinaryString gives the right outputs", func(t *testing.T) {

		cases := []struct {
			input    string
			expected string
		}{
			{"0", "0000"},
		}

		for _, c := range cases {
			binaryString := NewBinaryString(c.input)
			assert.Equal(t, c.expected, binaryString.bstring)
		}
	})

}
