package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	f, _ := os.Create("trace.out")
	defer f.Close()

	_ = trace.Start(f)

	ctx, task := trace.NewTask(context.Background(), "demo")
	defer task.End()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num string) {
			trace.WithRegion(ctx, num, func() {
				time.Sleep(5 * time.Second)
				fmt.Println(num)
			})
			wg.Done()
		}(fmt.Sprintf("region_%d", i))
	}
	wg.Wait()
	trace.Stop()
}
