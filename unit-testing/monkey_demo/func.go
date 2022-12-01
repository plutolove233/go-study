/*
@Coding : utf-8
@Time : 2022/4/16 17:07
@Author : 刘浩宇
@Software: GoLand
*/
package monkey_demo

import "monkey_demo/varys"

func MyFunc(uid int64) string{
	u,err := varys.GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}
	return "hello "+u.Name
}