package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// go tool pprof cpu.pprof查看性能消耗
// 可以通过list function_name 来查看具体函数分析

func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Println("recv from chan, value:", v)
		default:

		}
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool
	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Println("create cpu pprof failed, err: ", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)

	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Println("create mem pprof failed, err: ", err)
			return
		}
		pprof.WriteHeapProfile(file)
		defer file.Close()
	}
}
