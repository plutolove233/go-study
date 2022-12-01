package main

import "fmt"

type student struct{
	name string
	age int
}

func main(){
	m := make(map[string]*student)
	stus := []student{
		{
			name:"shyhao",
			age : 11,
		},
		{
			name:"pb",
			age:12,
		},
		{
			name:"pljj",
			age:13,
		},
	}
	for _,stu := range stus{
		m[stu.name] = &stu
		//fmt.Println(&stu)
		//fmt.Println(m["shyhao"])
	}

	for k,v := range m{
		fmt.Println(v)
		fmt.Println(k,"=>",v.name)
	}
}