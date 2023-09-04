package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryString(t *testing.T) {

	t.Run("Creating a NewBinaryString gives the right outputs", func(t *testing.T) {

		cases := []struct {
			input    string
			expected string
		}{
			{"0", "0000"},
			{"01", "0001"},
			{"101", "0101"},
			{"00", "0000"},
			{"10000", "00010000"},
		}

		for _, c := range cases {
			binaryString := NewBinaryString(c.input)
			assert.Equal(t, c.expected, binaryString.bstring)
		}
	})

    t.Run("Assert an input has a binary string format", func(t *testing.T) {
		cases := []struct {
			input    string
			expected bool
		}{
			{"0000", true},
			{"01", true},
			{"0300", false},
			{"0010", true},
			{"12345", false},
			{"-1", false},
			{"A&|.", false},
			{"", false},
		}
        for _, c := range cases {
            assert.Equal(t, c.expected, isBinaryFormat(c.input))
        }
    })
}
