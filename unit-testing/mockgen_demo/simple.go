/*
@Coding : utf-8
@Time : 2022/4/10 17:12
@Author : 刘浩宇
@Software: GoLand
*/
package mockgen_demo

type DB interface {
	Get(key string)(int,error)
	Add(key string, value int)error
}

func GetFromDB(db DB, key string) int {
	if v, err := db.Get(key);err != nil{
		return v
	}
	return -1
}