package main

import (
	"strings"
	"testing"
)

func TestCaptureInput(t *testing.T) {

	t.Run("Prints the input value back to user", func(t *testing.T) {

		input := "1010\n"
		rdr := strings.NewReader(input)

		captured, _ := captureInput(rdr)

		want := "1010"

		if captured != want {
			t.Error("unexpected result:", captured)
		}
	})

	t.Run("assert input is a binary string", func(t *testing.T) {

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
			rdr := strings.NewReader(c.input)
			captured, _ := captureInput(rdr)

		    ok := isBinaryString(captured)

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

func TestFindPQ(t *testing.T) {
	cases := []struct {
		m float64
		p float64
		q float64
	}{
		{5.0, 2.0, 3.0},
		{23.0, 4.0, 5.0},
		{102.0, 6.0, 7.0},
		{244.0, 7.0, 8.0},
		{456.0, 8.0, 9.0},
		{789.0, 9.0, 10.0},
		{1245.0, 10.0, 11.0},
		{3078.0, 11.0, 12.0},
		{5623.0, 12.0, 13.0},
		{9850.0, 13.0, 14.0},
		{15672.0, 13.0, 14.0},
	}
	t.Run("Find p and q given m", func(t *testing.T) {
		for _, c := range cases {
			p, q := findPandQ(c.m)
			if p != c.p {
				t.Errorf("got %f, want %f", p, c.p)
			}
			if q != c.q {
				t.Errorf("got %f, want %f", q, c.q)
			}
		}
	})
}

func TestDecimalToBinary(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{"3", "0011"},
        {"85", "01010101"},
		{"4", "0100"},
	}

	for _, c := range cases {

		b := decimalToBinary(c.input)
		if b != c.output {
			t.Errorf("got %q, want %q", b, c.output)
		}
	}

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
