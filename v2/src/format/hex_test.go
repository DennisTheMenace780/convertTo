package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexString(t *testing.T) {

	t.Run("Convert HexString to DecimalString", func(t *testing.T) {
		cases := []struct {
			input  Hex
			output Decimal
		}{
			{Hex("0"), Decimal("0")},
			{Hex("1"), Decimal("1")},
			{Hex("A"), Decimal("10")},
			{Hex("F"), Decimal("15")},
			{Hex("10"), Decimal("16")},
			{Hex("1F"), Decimal("31")},
			{Hex("100"), Decimal("256")},
			{Hex("3E8"), Decimal("1000")},
			{Hex("7FFF"), Decimal("32767")},
			{Hex("FFFF"), Decimal("65535")},
			{Hex("1FFFE"), Decimal("131070")},
			{Hex("7FFFFFFFFFFFFFFF"), Decimal("9223372036854775807")},
		}
		for _, c := range cases {
			decStr := c.input.ToDecimalString()
			assert.Equal(t, c.output, decStr)
		}
	})
}
