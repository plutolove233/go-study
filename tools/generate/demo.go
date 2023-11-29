/**
* @Author: yizhigopher
* @Description: go generate会执行go文件中的包含go:generate command的特殊注释，
*               当我们在控制台运行go generate时会执行这个command，但是不会执行源程序
* @File: demo.go
* @Version: 1.0.0
* @Date: 2023/11/29 19:19:31
 */

package demo

import "fmt"

//go:generate echo help
// 在执行go generata后会在控制台输出help，相当于执行echo help

func JustOneDemo() {
	fmt.Println("Hello go generate")
}
