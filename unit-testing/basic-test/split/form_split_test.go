/*
@Coding : utf-8
@Time : 2022/4/10 15:18
@Author : 刘浩宇
@Software: GoLand
*/
package split

import (
	"reflect"
	"testing"
)

func TestSplitAll(t *testing.T) {
	// define test form
	// we can use anonymous struct to define lots of example
	// and name all test example
	tests := []struct{
		name string
		input string
		sep string
		want []string
	}{
		{"base_case", "a:b:c",":",[]string{"a","b","c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}

	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Split(tt.input,tt.sep)
			if !reflect.DeepEqual(got,tt.want){
				t.Errorf("expected:%#v, got:%#v",tt.want,got)
			}
		})
	}
}