package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

var totalDuration time.Duration = 5

func getName(r io.Reader, w io.Writer) (string, error) {
	scanner := bufio.NewScanner(r)
	msg := "Your name please? Press the Enter key when Done"
	fmt.Fprintln(w, msg)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You entered an empty name")
	}
	return name, nil
}

func getNameContext(ctx context.Context) (string, error) {
	var err error
	name := "Default Name"
	c := make(chan error, 1)

	go func() {
		name, err = getName(os.Stdin, os.Stdout)
		c <- err
	}()

	select {
	case <-ctx.Done(): // ctx.Done() 함수가 반환하는 값을 쓰는 채널.
		return name, ctx.Err()
	case err := <-c:
		return name, err // getName() 함수가 반환할 때 쓰는 채널.
	}
}

func main() {
	allowedDuration := totalDuration * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), allowedDuration)
	defer cancel()

	name, err := getNameContext(ctx)

	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		fmt.Fprintln(os.Stdout, "%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, name)
}
