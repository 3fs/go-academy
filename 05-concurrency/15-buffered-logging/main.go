package main

import (
	"fmt"
	"time"
)

func worker(c <-chan string) {
	data := ""
	for {
		data += <-c
		if len(data) > 50 {
			// dump data do disk
			time.Sleep(5 * time.Second)
			data = ""
		}
	}
}

func httpHandler(c chan string) {
	start := time.Now()

	time.Sleep(100 * time.Millisecond)
	c <- "Log data"

	elapsed := time.Now().Sub(start)
	fmt.Println("Handler finished in", elapsed)
}

func main() {
	c := make(chan string, 5)
	for i := 0; i < 50; i++ {
		go worker(c)
	}

	for {
		go httpHandler(c)
		time.Sleep(200 * time.Millisecond)
	}
}
