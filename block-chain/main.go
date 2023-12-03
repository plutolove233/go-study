/**
* @Author: yizhigopher
* @Description: the learning of block chain, and it is cited from https://zhuanlan.zhihu.com/p/413585427
* @File: main.go
* @Version: 1.0.0
* @Date: 2023/12/03 20:03:35
 */

package main

import (
	"block/blockchain/pow"
	"fmt"
	"sync"
)

func main() {
	blockChain := pow.NewBlockChain()
	var wg sync.WaitGroup

	addFunc := func (i int) {
		blockChain.AddBlock(fmt.Sprintf("I am from %d", i))
		wg.Done()
	}
	
	wg.Add(3)
	for i := 1; i<=3; i++ {
		go addFunc(i)
	}
	wg.Wait()

	for _, b := range blockChain.Blocks {
		fmt.Printf("TimeStamp: %v\n", b.TimeStamp)
		fmt.Printf("data: %s\n", string(b.Data))
		fmt.Println("=====================")
	}
}
