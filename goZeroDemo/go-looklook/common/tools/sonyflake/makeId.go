// Package sonyflake
/*
@Coding : utf-8
@time : 2022/12/4 9:43
@Author : yizhigopher
@Software : GoLand
*/
package sonyflake

import (
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

var (
	sfOnce sync.Once
	sf     *Sonyflake
)

func GenerateID() uint64 {
	sfOnce.Do(func() {
		sf = NewSonyflake(Settings{
			StartTime: time.Now(),
		})
	})

	id, err := sf.NextID()
	if err != nil {
		logx.Errorf("generate UID failed\n")
	}
	return id
}
