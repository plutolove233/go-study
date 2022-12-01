/*
@Coding : utf-8
@Time : 2022/4/10 16:11
@Author : 刘浩宇
@Software: GoLand
*/
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlClient *gorm.DB
)

func InitMySQLClient() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"root",
		"localhost",
		"3306",
		"test",)
	mysqlClient, err  = gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		return err
	}
	return err
}

type Cars struct {
	Cost 	int 	`gorm:"column:cost;type:int(0)"`
	Id 		string	`gorm:"column:id;type:varchar(255)"`
	Brand 	string	`gorm:"column:brand;varchar(255)"`
}

func (Cars) TableName() string{
	return "cars"
}

func (m *Cars)recordStats() error{
	return mysqlClient.Create(&m).Error
}

func main() {
	err := InitMySQLClient()
	if err != nil {
		return
	}
	car := Cars{
		Cost:  100,
		Id:    "74LS138",
		Brand: "BMW",
	}
	err = car.recordStats()
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ok")
}