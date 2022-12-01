/*
@Coding : utf-8
@Time : 2022/4/10 14:45
@Author : 刘浩宇
@Software: GoLand
*/
package split

import (
	"fmt"
	"reflect"
	"testing"
)

/*
测试函数的基本要求：
	1. 函数名必须已Test开头
	2. 文件名必须已_test作为结尾
*/

/*
go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，
然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，
最后清理测试中生成的临时文件。
 */

/*
调用单元测试的命令：
go test 直接调用所有的测试函数
-v 显示详细信息
-run="name" 它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行
-bench=name 运行基准测试函数  为基准测试添加-benchmem参数，来获得内存分配的统计数据。
-run Example 运行示例函数
-cover 输出测试覆盖率
	-coverprofile=c.out 输出测试覆盖率到文件c.out
go tool cover -html=c.out 将c.out中测试覆盖率的内容已网页形式展示
 */

/*
测试函数	函数名前缀为Test		测试程序的一些逻辑行为是否正确
基准函数	函数名前缀为Benchmark	测试函数的性能
示例函数	函数名前缀为Example	为文档提供示例文档
*/

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func BenchmarkLoP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("shysyyh","s")
	}
}

func ExampleSplit() {
	//示例文档需要添加Output:注释行用来表明结果样式
	fmt.Println(Split("a:b:c",":"))
	fmt.Println(Split("shysyyh","s"))
	// Output:
	// [a b c]
	// [ hy yyh]
}