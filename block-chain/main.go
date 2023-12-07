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
	"block/transaction"
	"fmt"
)

func main() {
	txPool := make([]*transaction.Transaction, 0)
	var tempTx *transaction.Transaction
	var ok bool
	var property int
	chain := pow.NewBlockChain()
	property, _ = chain.FindUTXOs([]byte("Orion Liu"))
	fmt.Println("Balance of Orion Liu: ", property)

	tempTx, ok = chain.CreateTransaction([]byte("Orion Liu"), []byte("Krad"), 100)
	if ok {
		txPool = append(txPool, tempTx)
	}
	chain.AddBlock(txPool)
	txPool = make([]*transaction.Transaction, 0)
	property, _ = chain.FindUTXOs([]byte("Orion Liu"))
	fmt.Println("Balance of Orion Liu: ", property)

	tempTx, ok = chain.CreateTransaction([]byte("Krad"), []byte("Exia"), 200) // this transaction is invalid
	if ok {
		txPool = append(txPool, tempTx)
	}

	tempTx, ok = chain.CreateTransaction([]byte("Krad"), []byte("Exia"), 50)
	if ok {
		txPool = append(txPool, tempTx)
	}

	tempTx, ok = chain.CreateTransaction([]byte("Orion Liu"), []byte("Exia"), 100)
	if ok {
		txPool = append(txPool, tempTx)
	}
	chain.AddBlock(txPool)
	txPool = make([]*transaction.Transaction, 0)
	property, _ = chain.FindUTXOs([]byte("Orion Liu"))
	fmt.Println("Balance of Orion Liu: ", property)
	property, _ = chain.FindUTXOs([]byte("Krad"))
	fmt.Println("Balance of Krad: ", property)
	property, _ = chain.FindUTXOs([]byte("Exia"))
	fmt.Println("Balance of Exia: ", property)

	for _, block := range chain.Blocks {
		fmt.Printf("Timestamp: %v\n", block.TimeStamp)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("nonce: %d\n", block.Nonce)
		fmt.Println("Proof of Work validation:", block.ValidatePoW())
	}

	//I want to show the bug at this version.

	tempTx, ok = chain.CreateTransaction([]byte("Krad"), []byte("Exia"), 30)
	if ok {
		txPool = append(txPool, tempTx)
	}

	tempTx, ok = chain.CreateTransaction([]byte("Krad"), []byte("Orion Liu"), 30)
	if ok {
		txPool = append(txPool, tempTx)
	}

	chain.AddBlock(txPool)
	txPool = make([]*transaction.Transaction, 0)

	for _, block := range chain.Blocks {
		fmt.Printf("Timestamp: %v\n", block.TimeStamp)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("nonce: %d\n", block.Nonce)
		fmt.Println("Proof of Work validation:", block.ValidatePoW())
	}

	property, _ = chain.FindUTXOs([]byte("Orion Liu"))
	fmt.Println("Balance of Orion Liu: ", property)
	property, _ = chain.FindUTXOs([]byte("Krad"))
	fmt.Println("Balance of Krad: ", property)
	property, _ = chain.FindUTXOs([]byte("Exia"))
	fmt.Println("Balance of Exia: ", property)
}
