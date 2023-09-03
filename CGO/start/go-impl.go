package main

import "C"

import (
	"fmt"
)


//export SayHello
func SayHello(str *C.char){
	// implement C function by go
	fmt.Println("go impl", C.GoString(str))
}