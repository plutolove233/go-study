package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) reflect.Type{
	v := reflect.TypeOf(x)
	return v
}

func main(){
	fmt.Println("基本类型的类型反射")
	var a float32 = 3.14
	fmt.Println(a,"=>",reflectType(a))
	var b int64 = 313121
	fmt.Println(b,"=>",reflectType(b))

	fmt.Println("----------------------")
	fmt.Println("复杂类型反射")
	type Student struct {
		name string
		age int
	}
	type Book struct {
		name string
	}

	var d = Student{
		name: "saki",
		age:  12,
	}
	var B = Book{
		name:"Wuthering Heights",
	}
	fmt.Println("struct Student type=>",reflectType(d).Name(),reflectType(d).Kind())
	fmt.Println("struct Book type=>",reflectType(B).Name(),reflectType(B).Kind())
}