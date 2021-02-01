package main

import (
	"github.com/gin-gonic/gin"

	_ "seckill-server/docs"
	"seckill-server/model"
	"seckill-server/router"
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
	r := router.InitRouter()
	_ = r.Run(":7087") // listen and serve on 0.0.0.0:8080}
}
