/*
@Coding : utf-8
@Time : 2022/4/11 21:54
@Author : 刘浩宇
@Software: GoLand
*/
package GoStubStudy

import "io/ioutil"

var (
	configFile = "config.json"
	maxNum = 10
)


func GetConfig() ([]byte, error) {
	return ioutil.ReadFile(configFile)
}


func ShowNumber()int{
	// ...
	return maxNum
}