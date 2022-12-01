/*
@Coding : utf-8
@Time : 2022/4/16 17:11
@Author : 刘浩宇
@Software: GoLand
*/
package varys

type UserInfo struct {
	Name string
}

func GetInfoByUID(uid int64)(*UserInfo,error){
	if uid == 100 {
		return &UserInfo{
			Name: "shyhao",
		}, nil
	}
	return &UserInfo{}, nil
}