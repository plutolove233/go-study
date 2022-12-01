// Package controller
/*
@Coding : utf-8
@time : 2022/7/31 22:41
@Author : yizhigopher
@Software : GoLand
*/
package controller

import "github.com/gin-gonic/gin"

// GetMessage 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models._ResponsePostList
// @Router /posts2 [get]
func GetMessage(c *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	// 初始化结构体时指定初始参数
	c.JSON(200, gin.H{
		"code":    "2000",
		"message": "swagger test",
	})
}
