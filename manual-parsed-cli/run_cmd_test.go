package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		c      config
		input  string
		output string
		err    error
	}{
		{
			c:      config{numTimes: 3},
			input:  "",
			output: strings.Repeat("Your name please? Press the Enter key when done.\n", 1),
			err:    errors.New("You didn't enter your name"),
		},
		{
			c:      config{numTimes: 5},
			input:  "Guntak Kim",
			output: "Your name please? Press the Enter key when done.\n" + strings.Repeat("Nice to meet you Guntak Kim\n", 5),
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		rd := strings.NewReader(tc.input)
		err := runCmd(rd, byteBuf, tc.c)
		if err != nil && tc.err == nil {
			t.Fatalf("Expected error to be nil, but got %v", err)
		}
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error to be %v, but got %v", tc.err, err)
		}

		gotMsg := byteBuf.String()
		if gotMsg != tc.output {
			t.Errorf("Expected output to be %q, but got %q", tc.output, gotMsg)
		}
		byteBuf.Reset()
	}
}
