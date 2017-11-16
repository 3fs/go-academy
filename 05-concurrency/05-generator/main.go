package main

import (
	"fmt"
	"math/rand"
	"time"
)

func print(msg string) <-chan string { // returns receive-only channel of strings
	c := make(chan string)
	go func() { // launch the goroutine from inside the unction
		for i := 0; ; i++ {
			c <- fmt.Sprint(msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c // return the channel
}

func main() {
	c := print("Foo!")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}
