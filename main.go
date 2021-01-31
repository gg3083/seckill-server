package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"seckill-server/controller"
	_ "seckill-server/docs"
	"seckill-server/model"
	"seckill-server/setting"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example
// @license.name MIT
// @license.url https://Pay/blob/master/LICENSE
// @securityDefinitions.apikey Token
// @in header
// @name token
// @BasePath /
func main() {

	setting.InitSetting()
	model.InitMysql()

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	goodsApi := r.Group("/goods")
	{
		goodsApi.GET("/:id", controller.GetGoods)
		goodsApi.POST("/", controller.AddGoods)
	}

	userApi := r.Group("/user")
	{
		userApi.POST("/login", controller.Login)
		//初始化用户
		userApi.POST("/register", controller.Register)
		//添加收货地址
		userApi.POST("/address/")
	}
	_ = r.Run(":7087") // listen and serve on 0.0.0.0:8080}
}
