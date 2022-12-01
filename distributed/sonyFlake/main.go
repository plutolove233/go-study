package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Println("number of cpu: ", numCPU)

	var st Settings
	st.StartTime = time.Now()
	sf := NewSonyflake(st)
	if sf == nil {
		fmt.Println("sonyflake not created...")
		return
	}

	consumer := make(chan uint64)

	const numID = 10000

	generate := func() {
		for i := 0; i < numID; i++ {
			id, err := sf.NextID()
			if err != nil {
				fmt.Println("sonyflake next id generate failed...")
				return
			}
			consumer <- id
		}
	}

	for i := 0; i < 10; i++ {
		go generate()
	}

	set := make(map[uint64]struct{})
	for i := 0; i < 10*numID; i++ {
		id := <-consumer
		if _, ok := set[id]; ok {
			fmt.Println("duplicated id")
		}
		set[id] = struct{}{}
	}
	fmt.Println("number of id:", len(set))
}
