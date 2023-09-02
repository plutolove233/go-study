//hello.go
package main

//void SayHello(const char* s);
import "C"

func main(){
	C.SayHello(C.CString("Hello World\n"))
}

//SayHello函数放到当前目录下的一个C语言源文件中（后缀名必须是.c）同时需要用go run .来运行代码