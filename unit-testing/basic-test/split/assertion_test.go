/*
@Coding : utf-8
@Time : 2022/4/10 15:33
@Author : 刘浩宇
@Software: GoLand
*/
package split

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//断言的作用不需要再写大量的if-else判断语句，而且可以输出良好

func TestSomething(t *testing.T){
	assertions := assert.New(t)
	assertions.Equal(123,123,"they should be equal")
	assertions.NotEqual(123,456,"they should not be equal")
}