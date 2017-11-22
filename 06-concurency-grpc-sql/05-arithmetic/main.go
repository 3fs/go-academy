package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/3fs/go-academy/06-concurency-grpc-sql/05-artithmetic/arithmetic"
)

func main() {
	// initialize storage
	s := newStorage()
	ctx, cnc := context.WithCancel(context.Background())
	defer cnc()

	go readInput(s, cnc)

	// wait for the termination signal
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	select {
	case <-ctx.Done():
		fmt.Println("Context closed, stopping")
	case <-c:
		fmt.Println("Received a signal, stopping")
	}
}

func readInput(s *storage, cncFn func()) {
	for {
		// start the input line
		fmt.Print("\n> ")

		// read user's input
		r := bufio.NewReader(os.Stdin)
		input, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read user input: %v\n", err)
			cncFn()
		}

		// trim whitespace
		input = strings.TrimSpace(input)

		// calculate it
		result, err := arithmetic.Calculate(input, s)
		if err != nil {
			fmt.Printf("Error> %v\n", err)
			continue
		}

		// append to storage
		newPos := s.Append(result)

		// display the result
		fmt.Printf("$%d = %v\n", newPos, result)
	}
}
