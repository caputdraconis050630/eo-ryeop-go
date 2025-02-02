package main

import (
	"errors"
	"fmt"
	"github.com/caputdraconis050630/advanced-cli/sub-cmd-arch/cmd"
	"io"
	"os"
)

var errInvalidSubCommand = errors.New("Invalid sub-command specified")

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: mync [http|grpc] -h\n")
	cmd.HandleHttp(w, []string{"-h"})
	cmd.HandleGrpc(w, []string{"-h"})
}

func handleCommand(w io.Writer, args []string) error {
	var err error
	if len(args) < 1 {
		err = errInvalidSubCommand
	} else {
		switch args[0] {
		case "http":
			return cmd.HandleHttp(w, args[1:])
		case "grpc":
			return cmd.HandleGrpc(w, args[1:])
		case "-h":
			printUsage(w)
		case "-help":
			printUsage(w)
		default:
			err = errInvalidSubCommand
		}
	}

	if errors.Is(err, errInvalidSubCommand) || errors.Is(err, cmd.ErrNoServerSpecified) {
		fmt.Fprintln(w, err)
		printUsage(w)
	}
	return err
}

func main() {
	err := handleCommand(os.Stdout, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
