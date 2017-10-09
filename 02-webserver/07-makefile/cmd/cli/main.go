package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/3fs/go-academy/02-webserver/07-makefile/pkg/greeter"
)

func readStdin() string {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return ""
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		return ""
	}

	s, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return ""
	}

	return string(s)
}

func main() {
	var (
		name = flag.String("name", "", "Name of the person you'd like to greet")
	)

	// parse the flags
	flag.Parse()

	greet := readStdin()
	if greet == "" {
		greet = *name
	}

	fmt.Printf(greeter.Greet(greet))
}
