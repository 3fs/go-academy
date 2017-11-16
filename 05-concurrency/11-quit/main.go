package main

import (
	"fmt"
	"math/rand"
	"time"
)

func print(msg string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprint(msg, i):
				//do nothing
			case <-quit:
				return
			}
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	quit := make(chan bool)
	c := print("Foo!", quit)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	quit <- true
}
