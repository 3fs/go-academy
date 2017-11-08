package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx1, cnc1 := context.WithCancel(context.Background())
	ctx2, cnc2 := context.WithCancel(ctx1)
	ctx3, cnc3 := context.WithCancel(ctx2)
	ctx4, cnc4 := context.WithCancel(ctx3)

	// "use" cancel functions so the compiler doesn't complain
	_ = cnc1
	_ = cnc2
	_ = cnc3
	_ = cnc4

	go func() {
		for {
			select {
			case <-ctx1.Done():
				fmt.Println("Ctx1 done")
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("Ctx2 done")
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx3.Done():
				fmt.Println("Ctx3 done")
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx4.Done():
				fmt.Println("Ctx4 done")
				return
			}
		}
	}()

	cnc3()

	// make sure everything is done
	time.Sleep(1 * time.Second)
}
