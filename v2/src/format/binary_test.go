package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryString(t *testing.T) {

	t.Run(
		"Given an unformatted binary string, return a formatted binary string",
		func(t *testing.T) {

			cases := []struct {
				input    string
				expected Binary
			}{
				{"0", Binary("0000")},
				{"010", Binary("0010")},
				{"110", Binary("0110")},
				{"00001", Binary("00000001")},
			}

			for _, c := range cases {
				binaryString := NewBinaryString(c.input)
				assert.Equal(t, c.expected, binaryString)
			}
		},
	)

	t.Run("Assert an input has a binary string format", func(t *testing.T) {
		cases := []struct {
			input    string
			expected bool
		}{
			{"0000", true},
			{"01", true},
			{"0300", false},
			{"0010", true},
			{"10010", true},
			{"12345", false},
			{"-1", false},
			{"A&|.", false},
			{"", false},
		}
		for _, c := range cases {
			assert.Equal(t, c.expected, IsBinaryFormat(c.input))
		}
	})

	t.Run("Convert Binary to DecimalString", func(t *testing.T) {
		cases := []struct {
			input    string
			expected Decimal
		}{
			{"1010", Decimal("10")},
			{"10100", Decimal("20")},
			{"1000101", Decimal("69")},
			{"110100100", Decimal("420")},
			{"10101001100100111001", Decimal("694585")},
			{"10011101110101", Decimal("10101")},
			{"10111110111111", Decimal("12223")},
			{"0000", Decimal("0")},
		}

		for _, c := range cases {
			binaryString := NewBinaryString(c.input)
			decimalString := binaryString.ToDecimalString()
			assert.Equal(t, c.expected, decimalString)
		}
	})
}
