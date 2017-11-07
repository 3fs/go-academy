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
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	for i := range generate(ctx) {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
		if i == 5 {
			cancelFn()
		}
	}

	fmt.Println("done")
}
