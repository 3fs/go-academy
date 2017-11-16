package main

import (
	"fmt"
	"math/rand"
	"time"
)

func print(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprint(msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	c := fanIn(print("Foo!"), print("Bar!"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
