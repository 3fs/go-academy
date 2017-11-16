package main

import (
	"fmt"
	"math/rand"
	"time"
)

func print(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	go print("Foo!")
}
