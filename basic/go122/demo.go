package main

import "fmt"

type Interface interface {
	Do()
}

type A struct {
	name string
	age  int
}

type B struct {
	anaimal string
}

func (a *A) Do() {
	fmt.Println(a.name)
}

func (b *B) Do()  {
	fmt.Println(b.anaimal)
}

func Demo[T Interface](name T) {
	name.Do()
}

func (t *T)()  {
	
}

func main(){
	a := &A{
		age: 12,
		name: "A",
	}
	// 可以从赋值推断类型参数
	Demo(a)
}