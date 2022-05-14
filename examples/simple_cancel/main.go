package main

import (
	"context"
	"fmt"
	"time"
)

// The main idea of context is that parent task can cancel ongoiong child tasks
// Suppose you are going to the airport, if you received a msg from the airline that the flight is cancled,
// you should be able to react on that and not continue going to the airport anymore
// In programming, context is used to control this.

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	sleepAndPrint(ctx, time.Second*5, "hello")
}

func sleepAndPrint(ctx context.Context, d time.Duration, msg string) {
	select {
	case n := <-ctx.Done():
		fmt.Println("ctx is done", n)
	case <-time.After(d):
		fmt.Println(msg)
	}
}
