package main

import (
	"strings"
	"testing"
)

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

		cases := []struct {
			input    string
			expected bool
		}{
			{"0000", true},
			{"01", false},
			{"0300", false},
			{"0010", true},
			{"12345", false},
			{"-1", false},
			{"A&|.", false},
			{"", false},
		}
		for _, c := range cases {
			rdr := strings.NewReader(c.input)
			captured := captureInput(rdr)

			_, ok := is4DigitBinaryString(captured)

			if ok != c.expected {
				t.Error("Detected an invalid input: ", captured)
			}

		}

	})

}

func TestConvert(t *testing.T) {
	t.Run("binary string to decimal representation", func(t *testing.T) {

		cases := []struct {
			input    string
			expected string
		}{
			{"00", "0"},
			{"01", "1"},
			{"10", "2"},
			{"11", "3"},
			{"0000", "0"},
			{"0100", "4"},
			{"1001", "9"},
			{"1010101", "85"},
		}

		for _, c := range cases {
			val := binaryToDecimal(c.input)
			if val != c.expected {
				t.Errorf("got %q, want %q", val, c.expected)
			}
		}
	})
}

// func TestCaptureInputPipe(t *testing.T) {
// 	r, w, _ := os.Pipe()
// 	w.Write([]byte("1010\n"))
//
// 	got := captureInput2(r)
// 	want := "1010"
// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }

// func TestCaptureInput2(t *testing.T) {
//
// 	t.Run("Prints the input value back to user", func(t *testing.T) {
//
// 		want := "1010"
//
// 		// Create a temporary file for IO
// 		tmpfile, err := os.CreateTemp("", "exampleDaddy")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
//
// 		// Clean up file
// 		defer os.Remove(tmpfile.Name())
//
// 		// Write to the tmpfile
// 		if _, err := tmpfile.WriteString(want); err != nil {
// 			log.Fatal(err)
// 		}
//
// 		// Seek sets the offset for the next Read or Write on file to offset, interpreted
// 		// according to whence: 0 means relative to the origin of the file, 1 means relative
// 		// to the current offset, and 2 means relative to the end.
// 		if _, err := tmpfile.Seek(0, 0); err != nil {
// 			log.Fatal(err)
// 		}
//
// 		// assign the old value of os.Stdin to something and restore later
// 		oldStdin := os.Stdin
// 		defer func() { os.Stdin = oldStdin }()
//
// 		// Set the stdin to what was written into tmpfile
// 		os.Stdin = tmpfile
//
// 		captured := captureInput2()
//
// 		if captured != want {
// 			t.Error("unexpected result:", want)
// 		}
// 	})
// }

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
