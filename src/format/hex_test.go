package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexString(t *testing.T) {

	t.Run("Convert HexString to DecimalString", func(t *testing.T) {
		cases := []struct {
			input  string
			output string
		}{
			{"0", "0"},
			{"1", "1"},
			{"A", "10"},
			{"F", "15"},
			{"10", "16"},
			{"1F", "31"},
			{"100", "256"},
			{"3E8", "1000"},
			{"7FFF", "32767"},
			{"FFFF", "65535"},
			{"1FFFE", "131070"},
			{"7FFFFFFFFFFFFFFF", "9223372036854775807"},
		}
		for _, c := range cases {
			hexStr := NewHexString(c.input)
			decStr := hexStr.ToDecimalString()
			assert.Equal(t, c.output, decStr.Dstring)
		}
	})
}
