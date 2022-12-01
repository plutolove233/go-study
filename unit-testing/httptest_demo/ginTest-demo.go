/*
@Coding : utf-8
@Time : 2022/4/10 15:40
@Author : 刘浩宇
@Software: GoLand
*/
package httptest_demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Parser struct {
	Name string	`json:"Name"`
}

func helloHandler(c *gin.Context) {
	var parser Parser
	err := c.ShouldBindJSON(&parser)
	if err != nil {
		c.JSON(200,gin.H{
			"code":"4004",
			"message":"获取请求信息失败",
		})
		return
	}

	c.JSON(200,gin.H{
		"code":"2000",
		"message":fmt.Sprintf("Hello %s",parser.Name),
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/hello",helloHandler)
	return router
}