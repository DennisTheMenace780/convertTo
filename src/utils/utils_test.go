package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {

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
			chunkedString := Chunk(c.input, c.chunks)
			assert.Equal(t, c.expected, chunkedString)
		}
	})
}
