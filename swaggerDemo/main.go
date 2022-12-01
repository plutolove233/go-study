package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"swaggerDemo/controller"
	_ "swaggerDemo/docs"
)

// @title swagger测试学习
// @version 1.0
// @description swagger生成接口文档测试
// @termsOfService http://swagger.io/terms

// @contact.name yizhigopher
// @contact.url http://swagger.io/support
// @contact.email yizhigopher@foxmail.com

// @host 0.0.0.0
// @BasePath .

func main() {
	r := gin.New()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.GET("/test", controller.GetMessage)
	r.Run()
}
