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

	t.Run("Convert binary to DecimalString", func(t *testing.T) {
		cases := []struct {
			input    string
			expected string
		}{
			{"1010", "10"},
			{"10100", "20"},
			{"1000101", "69"},
			{"110100100", "420"},
			{"10101001100100111001", "694585"},
			{"10011101110101", "10101"},
			{"10111110111111", "12223"},
			{"0000", "0"},
		}

		for _, c := range cases {
			binaryString := NewBinaryString(c.input)
			decimalString := binaryString.ToDecimalString()
			assert.Equal(t, c.expected, decimalString.dstring)
		}
	})

	t.Run("Chunk string chunks properly", func(t *testing.T) {
		cases := []struct {
			input    string
			chunks   int
			expected []string
		}{
			{"10101010", 1, []string{"1", "0", "1", "0", "1", "0", "1", "0"}},
			{"10101010", 2, []string{"10", "10", "10", "10"}},
			{"10101010", 4, []string{"1010", "1010"}},
			{"10101010", 5, []string{"10101", "010"}},
			{"10101010", 8, []string{"10101010"}},
		}

		for _, c := range cases {
			binaryString := NewBinaryString(c.input)
			chunkedString := binaryString.chunkString(c.chunks)
			assert.Equal(t, c.expected, chunkedString)
		}
	})
}
