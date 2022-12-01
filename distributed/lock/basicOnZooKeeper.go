package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second)
	if err != nil {
		panic(err)
	}

	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	err = l.Lock()
	if err != nil {
		panic(err)
	}

	println("lock success, do your bussiness logic")

	time.Sleep(time.Second * 2)

	l.Unlock()

	println("unlock success, finish business logic")
}
