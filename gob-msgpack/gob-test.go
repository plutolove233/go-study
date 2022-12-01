/*
@Coding : utf-8
@Time : 2022/4/9 15:43
@Author : 刘浩宇
@Software: GoLand
*/
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type s struct {
	data map[string]interface{}
}

func main() {
	// gobDemo
	s1 := s{
		data:make(map[string]interface{},8),
	}
	s1.data["count"] = 1
	//encode
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	err := enc.Encode(s1.data)
	if err != nil {
		fmt.Println("gob encode failed,err:",err.Error())
		return
	}

	b := buff.Bytes()
	fmt.Println(b)

	s2 := s{
		data: make(map[string]interface{},8),
	}
	//decode
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	if err = dec.Decode(&s2.data); err!=nil{
		fmt.Println("gob decode failed,err:",err.Error())
		return
	}
	fmt.Println(s2.data)
}