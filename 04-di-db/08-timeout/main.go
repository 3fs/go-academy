package main

import (
	"context"
	"fmt"
	"time"
)

func generate(ctx context.Context) chan int {
	ch := make(chan int)
	n := 1

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("for done")
				close(ch)
				return
			case ch <- n:
				n++
			}
		}
	}()

	return ch
}

func main() {
	ctx, cancelFn := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFn()

	for i := range generate(ctx) {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("done")
}
