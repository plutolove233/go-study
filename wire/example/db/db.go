// Package db
/*
@Coding : utf-8
@Time : 2023/2/9 22:05
@Author : yizhigopher
@Software : GoLand
*/
package db

import "database/sql"

type Config struct {
}

func NewDB(c *Config) (*sql.DB, error) {
	return &sql.DB{}, nil
}
