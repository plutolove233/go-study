/*
@Coding : utf-8
@Time : 2022/4/9 15:14
@Author : 刘浩宇
@Software: GoLand
*/
package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name","shyhao","姓名")
	age := flag.Int("age",12,"年龄")
	married := flag.Bool("married",false,"婚否")
	delay := flag.Duration("d",0,"时间间隔")

	//example
	flag.Parse()
	fmt.Println(*name,*age,*married,*delay)
	fmt.Println(flag.Args())
	fmt.Println(flag.NFlag())
	fmt.Println(flag.NArg())
}