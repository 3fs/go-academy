package main

import "fmt"

func count(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int, 5)
	go count(c)

	for i := range c {
		fmt.Println(i)
	}
}
