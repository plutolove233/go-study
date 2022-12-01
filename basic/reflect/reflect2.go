package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	t := reflect.ValueOf(x)
	k := t.Kind()
	switch k {
	case reflect.Int:
		fmt.Println("type is int, value is ", t.Int())
	default:
		fmt.Println(t)
	}
}

func main() {
	reflectValue(12)
	reflectValue(map[string]interface{}{
		"code": 2000,
	})
}
