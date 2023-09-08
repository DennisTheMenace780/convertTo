package format

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
			assert.Equal(t, c.expected, binaryString.Bstring)
		}
	})

	t.Run("Converts decimal string to binary string", func(t *testing.T) {
		cases := []struct {
			input    string
			expected string
		}{
			{"0", "0"},
		}
        for _, c := range cases {
            dstring := NewDecimalString(c.input) 
            hexStr := dstring.ToBinaryString()
            assert.Equal(t, c.expected, hexStr.Bstring)
        }
	})

	t.Run("Converts decimal string to Hex string", func(t *testing.T) {
		cases := []struct {
			input    string
			expected string
		}{
			{"0", "0"},
			{"1", "1"},
			{"10", "A"},
			{"11", "B"},
			{"12", "C"},
			{"13", "D"},
			{"14", "E"},
			{"15", "F"},
			{"16", "10"},
			{"456", "1C8"},
			{"31148", "79AC"},
		}
        for _, c := range cases {
            dstring := NewDecimalString(c.input) 
            hexStr := dstring.ToHexString()
            assert.Equal(t, c.expected, hexStr.hstring)
        }
	})
}
