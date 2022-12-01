package main

import "fmt"

func boilingExample() {
	const boilingF = 212.2
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g\n", c)
}

func variableExample() {
	var i, j, k int
	var b, f, s = true, 2.3, "four"
	fmt.Println(i, j, k, b, f, s)
}

func pointerExample() {
	x := 1
	p := &x
	fmt.Println(*p, p)

}

func main() {
	boilingExample()
	variableExample()
	pointerExample()
}
