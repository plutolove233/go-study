package main

import "sync"

// generics first example
type Lockable [T any] struct {
	mut sync.Mutex
	Data T
}

func main(){

	var n Lockable[int]
	n.mut.Lock()
	n.Data++
	n.mut.Unlock()

	var f Lockable[float32]
	f.mut.Lock()
	f.Data+=1.23
	f.mut.Unlock()

	var b Lockable[[]byte]
	b.mut.Lock()
	b.Data = []byte("12312")
	b.mut.Unlock()

	res := max(1, 'A')
	println(res)
}
/*
从上述例子中，Lockable是一种泛型类型。对比非泛型类型，它有个与众不同的地方，一个类型参数列表，
在这个泛型声明中。它的类型形参列表的是[T any]

一个类型形参可能会包含一个或多个用方括号包含，且用竖线分隔的类型形参声明。
每个参数声明包含一个类型参数名称和类型约束。例如，T是类型参数的名称，any是T的类型约束。

请注意any是一种在go1.18中新增的预声明标识符。它是空接口interface{}的别称。我们需要明白所有类型均实现了空接口。

我们可以将约束视为类型的类型（类型参数）。所有的约束都是一种接口类型。约束时泛型的核心，并会在下一章详细阐明
*/