package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

// TODO: Analyze test cases to make sure they're doing what they should be doing. 
// Am I testing everything correctly?
// Am I covering else branches?

func TestCaptureInput(t *testing.T) {

	t.Run("Prints the input value back to user", func(t *testing.T) {

		input := "1010\n"
		rdr := strings.NewReader(input)

		captured := captureInput(rdr)
		want := "1010"

		if captured != want {
			t.Error("unexpected result:", captured)
		}
	})

	t.Run("assert input is a 4 digit binary string", func(t *testing.T) {

		input := "1010\n"
		rdr := strings.NewReader(input)

		captured := captureInput(rdr)
		length, ok := is4DigitBinaryString(captured)

		expectedLength := 4

		if !ok {
			t.Error("Detected a non-binary string: ", captured)
		}

		if length != expectedLength {
			t.Errorf("Got length of %q, want length of %q", length, expectedLength)
		}

	})
}

func TestCaptureInput2(t *testing.T) {

	t.Run("Prints the input value back to user", func(t *testing.T) {

		want := "1010"

		// Create a temporary file for IO
		tmpfile, err := os.CreateTemp("", "exampleDaddy")
		if err != nil {
			log.Fatal(err)
		}

		// Clean up file
		defer os.Remove(tmpfile.Name())

		// Write to the tmpfile
		if _, err := tmpfile.WriteString(want); err != nil {
			log.Fatal(err)
		}

		// Seek sets the offset for the next Read or Write on file to offset, interpreted
		// according to whence: 0 means relative to the origin of the file, 1 means relative
		// to the current offset, and 2 means relative to the end.
		if _, err := tmpfile.Seek(0, 0); err != nil {
			log.Fatal(err)
		}

		// assign the old value of os.Stdin to something and restore later
		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		// Set the stdin to what was written into tmpfile
		os.Stdin = tmpfile

		captured := captureInput2()

		if captured != want {
			t.Error("unexpected result:", want)
		}
	})
}

// func TestCaptureMultipleInputs(t *testing.T) {
//
// 	cases := []struct {
// 		input    string
// 		expected string
// 	}{
// 		{"0101", "0101"},
// 		{"1010", "1010"},
// 		{"0000", "0000"},
// 	}
//
// 	for _, c := range cases {
// 		t.Run(fmt.Sprintf("Echos the input %s back to user", c.input), func(t *testing.T) {
//
// 			rdr := strings.NewReader(c.input)
//
// 			capturedInput := captureInput(rdr)
// 			want := c.expected
//
// 			if capturedInput != want {
// 				t.Error("unexpected result:", c.input)
// 			}
// 		})
// 	}
// }
