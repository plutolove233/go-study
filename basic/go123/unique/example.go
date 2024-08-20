package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"unique"
)

// 字符串驻留是一种存储技术，确保每个独特的字符串值在内存中只存在一个副本。
// 这意味着如果两个字符串具有相同的值，它们将共享同一个内存地址，而不是每个字符串都占用独立的内存空间。

// example

func wordGen(nDistinct, wordLen int) func() string {
	vocab := make([]string, nDistinct)
	for i := range nDistinct {
		word := randomString(wordLen)
		vocab[i] = word
	}
	return func() string {
		word := vocab[rand.Intn(nDistinct)]
		return word
	}
}

func randomString(n int) string {
	const letters = "eddycjyabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		b := make([]byte, 1)
		if _, err := rand.Read(b); err != nil {
			panic(err)
		}
		ret[i] = letters[int(b[0])%len(letters)]
	}
	return string(ret)
}

var words []string
var wordsWithUnique []unique.Handle[string]

func main() {
	const nWords = 10000
	const nDistinct = 100
	const wordLen = 40
	generate := wordGen(nDistinct, wordLen)
	memBefore := getAlloc()

	words = make([]string, nWords)
	for i := range nWords {
		words[i] = generate()
	}

	memAfter := getAlloc()
	memUsed := memAfter - memBefore
	fmt.Println("======== without unqiue ========")
	fmt.Printf("Memory used: %dKB\n", memUsed/1024)

	generate = wordGen(nDistinct, wordLen)
	memBefore = getAlloc()

	wordsWithUnique = make([]unique.Handle[string], nWords)
	for i := range nWords {
		wordsWithUnique[i] = unique.Make(generate())
	}

	memAfter = getAlloc()
	memUsed = memAfter - memBefore
	fmt.Println("======== with unqiue ========")
	fmt.Printf("Memory used: %dKB\n", memUsed/1024)
}

func getAlloc() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}
