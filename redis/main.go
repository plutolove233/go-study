package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

var (
	rdb *redis.Client
)

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.170.129:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func V8Example() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		return
	}

	err := rdb.SetEX(ctx, "key", "value", 1*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func ZSetExample() {
	zsetKey := "language_rank"

	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 99.0, Member: "C/C++"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	if err := rdb.ZAdd(ctx, zsetKey, languages...).Err(); err != nil {
		fmt.Println("zadd failed, err:", err)
		return
	}

	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Println("zincrby failed, err:", err)
		return
	}
	fmt.Println("Golang's score is ", newScore)

	ret := rdb.ZRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Println("z range by score failed, err: ", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Println("------------------", err)
		return
	}
	V8Example()
	ZSetExample()
}
