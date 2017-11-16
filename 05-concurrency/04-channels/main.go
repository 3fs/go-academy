package main

import (
	"fmt"
	"math/rand"
	"time"
)

func print(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprint(msg, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go print("Foo!", c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}
