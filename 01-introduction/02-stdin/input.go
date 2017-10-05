package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, _ := readStdin()
	fmt.Printf("Hello %s!\n", in)
}

func readStdin() (string, error) {
	b, _, err := bufio.NewReader(os.Stdin).ReadLine()
	return string(b), err
}
