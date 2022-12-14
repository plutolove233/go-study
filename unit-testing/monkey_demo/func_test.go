/*
@Coding : utf-8
@Time : 2022/4/16 17:09
@Author : 刘浩宇
@Software: GoLand
*/
package monkey_demo

import (
	"bou.ke/monkey"
	"monkey_demo/varys"
	"strings"
	"testing"
)

func TestMyFunc(t *testing.T) {
	// 对 varys.GetInfoByUID 进行打桩
	// 无论传入的uid是多少，都返回 &varys.UserInfo{Name: "liwenzhou"}, nil
	monkey.Patch(varys.GetInfoByUID, func(int64)(*varys.UserInfo, error) {
		return &varys.UserInfo{Name: "liwenzhou"}, nil
	})

	ret := MyFunc(123)
	if !strings.Contains(ret, "liwenzhou"){
		t.Fatal()
	}
}