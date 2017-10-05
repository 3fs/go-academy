package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "", "Name of the person you’d like to greet")
	flag.Parse()

	fmt.Printf("Hello %s!\n", *name)
}
