package main

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{"ValidName", "John\n", "John", false},
		{"EmptyName", "\n", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			w := &bytes.Buffer{}

			result, err := getName(r, w)

			if (err != nil) != tt.hasError {
				t.Errorf("getName() error = %v, wantErr %v", err, tt.hasError)
				return
			}
			if result != tt.expected {
				t.Errorf("getName() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetNameContext(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		timeout     time.Duration
		expected    string
		expectedErr error
	}{
		{"NameBeforeTimeout", "Alice\n", 2 * time.Second, "Alice", nil},
		{"TimeoutBeforeName", "", 1 * time.Millisecond, "Default Name", context.DeadlineExceeded},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			w := &bytes.Buffer{}

			oldStdin := os.Stdin
			oldStdout := os.Stdout
			defer func() { os.Stdin, os.Stdout = oldStdin, oldStdout }()

			os.Stdin = r
			os.Stdout = w

			ctx, cancel := context.WithTimeout(context.Background(), tt.timeout)
			defer cancel()

			result, err := getNameContext(ctx)

			if err != tt.expectedErr {
				t.Errorf("getNameContext() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}
			if result != tt.expected {
				t.Errorf("getNameContext() = %v, want %v", result, tt.expected)
			}
		})
	}
}
