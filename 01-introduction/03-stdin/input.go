package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	in, _ := readStdin()
	fmt.Printf("Hello %s!\n", in)
}

func readStdin() (string, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		return "", errors.New("StdIn not a named pipe")
	}

	b, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return "", err
	}

	return string(b), nil
}
