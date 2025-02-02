package cmd

import (
	"bytes"
	"errors"
	"testing"
)

func TestHandleGrpc(t *testing.T) {
	usageMessage := `
grpc: A gRPC Client

grpc: <options> server

Options: 
  -body string
    	Body of request
  -method string
    	Method to call
`
	testConfigs := []struct {
		args   []string
		output string
		err    error
	}{
		{
			// 위치 인수를 지정하지 않고 서브커맨드로 grpc를 호출하는 동작 테스트
			args: []string{},
			err:  ErrNoServerSpecified,
		},
		{
			// 인수로 "-h"를 지정하고 서브커맨드 grpc를 호출하는 동작 테스트
			args:   []string{"-h"},
			err:    errors.New("flag: help requested"),
			output: usageMessage,
		},
		{
			// 위치 인수로 서버의 URL을 지정하고 서브커맨드 grpc를 호출하는 동작 테스트
			args:   []string{"-method", "service.host.local/method", "-body", "{}", "http://localhost"},
			err:    nil,
			output: "Executing grpc command\n",
		},
	}
	byteBuf := new(bytes.Buffer)
	for _, tc := range testConfigs {
		err := HandleGrpc(byteBuf, tc.args)
		if tc.err == nil && err != nil {
			t.Fatalf("Expected nil error, got %v", err)
		}

		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error %v, got %v", tc.err, err)
		}

		if len(tc.output) != 0 {
			gotOutput := byteBuf.String()
			if tc.output != gotOutput {
				t.Errorf("Expected output to be: %#v, Got: %#v", tc.output, gotOutput)
			}
		}
		byteBuf.Reset()
	}
}
