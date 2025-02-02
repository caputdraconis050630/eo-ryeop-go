package main

import (
	"bytes"
	"testing"
)

func TestHandleCommand(t *testing.T) {
	usageMessage := "Usage: mync [http|grpc] -h\n\nhttp: A HTTP Client\n\nhttp: <options> server\n\nOptions: \n  -verb string\n    \tHTTP method (default \"GET\")\n\ngrpc: A gRPC Client\n\ngrpc: <options> server\n\nOptions: \n  -body string\n    \tBody of request\n  -method string\n    \tMethod to call\n"

	testConfigs := []struct {
		args   []string
		output string
		err    error
	}{
		{
			// 애플리케이션에 인수가 저장되지 않은 경우의 동작 테스트
			args:   []string{},
			err:    errInvalidSubCommand,
			output: "Invalid sub-command specified\n" + usageMessage,
		},
		{
			// 애플리케이션에 인수로 -h가 지정된 경우 동작 테스트
			args:   []string{"-h"},
			err:    nil,
			output: usageMessage,
		},
		{
			// 애플리케이션에 인수로 -h가 지정된 경우 동작 테스트
			args:   []string{"foo"},
			err:    errInvalidSubCommand,
			output: "Invalid sub-command specified\n" + usageMessage,
		},
	}

	byteBuf := new(bytes.Buffer)

	for _, tc := range testConfigs {
		err := handleCommand(byteBuf, tc.args)
		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, got %v", err)
		}

		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error %v, got %v", tc.err, err)
		}

		if len(tc.output) != 0 {
			gotOutput := byteBuf.String()
			if tc.output != gotOutput {
				t.Errorf("Expected output to be: %#v, Got %#v", tc.output, gotOutput)
			}
		}

		byteBuf.Reset()
	}
}
