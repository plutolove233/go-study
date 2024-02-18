package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	merged := append(a, b...)
	fmt.Println(merged)

	m2 := slices.Concat[[]int](a, b)
	fmt.Println(m2)
}