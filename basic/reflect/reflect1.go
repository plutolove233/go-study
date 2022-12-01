package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v.Name(), v.Kind())
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int = 4
	reflectType(b)
}
