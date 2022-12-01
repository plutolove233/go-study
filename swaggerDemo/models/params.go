// Package models
/*
@Coding : utf-8
@time : 2022/7/31 22:47
@Author : yizhigopher
@Software : GoLand
*/
package models

type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}

type _ResponsePostList struct {
	Code    int    `json:"code"`    // 业务响应状态码
	Message string `json:"message"` // 提示信息
}
