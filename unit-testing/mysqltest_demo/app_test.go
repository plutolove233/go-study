/*
@Coding : utf-8
@Time : 2022/4/10 16:43
@Author : 刘浩宇
@Software: GoLand
*/
package main

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

//waiting
//question:how to mock mysql database with the usage of "gorm.io"

func TestShouldUpdateStats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected",err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE cars").WillReturnResult(sqlmock.NewResult(1,1))
}