package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(50*time.Millisecond))

	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("over slept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
