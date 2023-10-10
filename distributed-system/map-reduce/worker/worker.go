package main

import "distributed-system/map-reduce/utils"

func main() {
	println("worker node is running...")
	Worker(utils.Map, utils.Reduce)
}
