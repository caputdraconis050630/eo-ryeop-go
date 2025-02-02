package main

import (
	"bytes"
	"context"
	"io"
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
	// 표준 입력 리다이렉션
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r
	go func() {
		w.WriteString("TestInput\n")
		w.Close()
	}()

	// 표준 출력 캡처
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	outR, outW, _ := os.Pipe()
	os.Stdout = outW

	// 테스트 실행 및 결과 확인
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result, _ := getNameContext(ctx)

	outW.Close()
	capturedOut, _ := io.ReadAll(outR)

	if result != "TestInput" {
		t.Errorf("Expected TestInput, got %s", result)
	}
	if !strings.Contains(string(capturedOut), "Your name please?") {
		t.Error("Missing prompt message")
	}
}
