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

func main() {
	c := print("Foo!")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(800 * time.Millisecond):
			fmt.Println("too slow")
			return
		}
	}
}
