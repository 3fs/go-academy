package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/3fs/go-academy/03-log-vendor/03-middleware/pkg/greeter"
)

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

func main() {
	var (
		name = flag.String("name", "", "Name of the person you'd like to greet")
	)

	// parse the flags
	flag.Parse()

	greet, _ := readStdin()
	if greet == "" {
		greet = *name
	}

	var service greeter.Service
	{
		service = greeter.New()
		service = greeter.ValidateMiddleware()(service)
	}

	msg := service.Greet(greet)
	if msg == "" {
		fmt.Println("Missing name")
		os.Exit(1)
	}

	fmt.Println(msg)
}
