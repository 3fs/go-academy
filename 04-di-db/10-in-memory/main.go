package main

import (
	"fmt"
	"time"

	"github.com/3fs/go-academy/04-di-db/10-in-memory/kvMemory"
)

func main() {
	kv := kvMemory.New()
	kv.Add("first", "value")
	kv.Add("second", "value")

	// uncomment for race condition
	go kv.Add("first", "value")

	time.Sleep(50 * time.Millisecond)

	v, err := kv.Get("second")
	fmt.Printf("Second value = %s; err = %v\n", v, err)

	_ = kv.Remove("second")
	v, err = kv.Get("Second")
	fmt.Printf("Second value = %s; err = %v\n", v, err)
}
