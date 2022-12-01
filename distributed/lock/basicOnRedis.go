package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

func incr() {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	var (
		lockKey    = "counter_lock"
		counterKey = "counter"
	)

	resp := client.SetNX(ctx, lockKey, 1, time.Second*5)
	lockSuccess, err := resp.Result()

	if err != nil || !lockSuccess {
		fmt.Println(err, " lock result: ", lockSuccess)
		return
	}

	getResp := client.Get(ctx, counterKey)
	cntValue, err := getResp.Int64()
	if err == nil {
		cntValue++
		resp := client.Set(ctx, counterKey, cntValue, 0)
		_, err := resp.Result()
		if err != nil {
			println("set counter value error, ", err)
		}

		println("current counter is ", cntValue)
		delResp := client.Del(ctx, lockKey)
		unlockSuccess, err := delResp.Result()
		if err == nil && unlockSuccess > 0 {
			println("unlock success!")
		} else {
			println("unlock failed, err=", err)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
