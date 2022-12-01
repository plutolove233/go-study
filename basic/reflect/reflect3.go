package main

import (
	"fmt"
	"reflect"
)

type Base struct {
	Basis  int `json:"basis"`
	BasisB int `json:"basis-b"`
}

type testParser struct {
	X Base `json:"x"`
	A int  `json:"a"`
	B int  `json:"b"`
}

func StructToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			temp := v.Field(i).Interface()
			if !reflect.DeepEqual(temp, reflect.Zero(reflect.TypeOf(temp)).Interface()) {
				out[tagValue] = temp
			}
		}
	}
	return out, nil
}

func MapToStruct(in map[string]interface{}, out interface{}) error {
	v := reflect.ValueOf(out)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()

	for key, value := range in {
		_, ok := t.FieldByName(key)
		if ok {
			v.FieldByName(key).Set(reflect.ValueOf(value))
		}
	}

	return nil
}

func main() {
	l := testParser{
		X: Base{
			Basis:  3,
			BasisB: 4,
		},
		A: 1,
		B: 2,
	}
	m, e := StructToMap(l, "json")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(m)

	t := testParser{}
	if err := MapToStruct(map[string]interface{}{
		"X": Base{
			Basis:  1,
			BasisB: 2,
		},
		"A": 3,
		"B": 4,
	}, &t); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(t)
}
