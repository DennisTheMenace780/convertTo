package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecimalString(t *testing.T) {

	t.Run("Converts Decimal to Binary", func(t *testing.T) {
		cases := []struct {
			input    Decimal
			expected Binary
		}{
			{Decimal("10"), Binary("1010")},
			{Decimal("20"), Binary("10100")},
			{Decimal("69"), Binary("1000101")},
			{Decimal("420"), Binary("110100100")},
			{Decimal("694585"), Binary("10101001100100111001")},
			{Decimal("10101"), Binary("10011101110101")},
			{Decimal("12223"), Binary("10111110111111")},
			{Decimal("0"), Binary("0")},
		}
		for _, c := range cases {
			decStr := c.input
			binaryString := decStr.ToBinaryString()
			assert.Equal(t, c.expected, binaryString)
		}
	})

	// t.Run("Converts decimal string to Hex string", func(t *testing.T) {
	// 	cases := []struct {
	// 		input    string
	// 		expected string
	// 	}{
	// 		{"0", "0"},
	// 		{"1", "1"},
	// 		{"10", "A"},
	// 		{"11", "B"},
	// 		{"12", "C"},
	// 		{"13", "D"},
	// 		{"14", "E"},
	// 		{"15", "F"},
	// 		{"16", "10"},
	// 		{"456", "1C8"},
	// 		{"31148", "79AC"},
	// 	}
	//        for _, c := range cases {
	//            dstring := NewDecimalString(c.input)
	//            hexStr := dstring.ToHexString()
	//            assert.Equal(t, c.expected, hexStr.hstring)
	//        }
	// })
}
