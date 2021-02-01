package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"seckill-server/pkg/util"
	"seckill-server/router/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	goodsApi := r.Group("/goods")
	{
		goodsApi.GET("/:id", v1.GetGoods)
	}

	goodsAuthApi := r.Group("/auth/goods")
	goodsAuthApi.Use(util.JWT())
	{
		goodsAuthApi.POST("/", v1.AddGoods)

	}

	userApi := r.Group("/user")
	{
		userApi.POST("/login", v1.Login)
		//初始化用户
		userApi.POST("/register", v1.Register)

	}

	userAuthApi := r.Group("/auth/user")
	userAuthApi.Use(util.JWT())
	{
		//添加收货地址
		userAuthApi.POST("/address", v1.Address)
	}

	orderApi := r.Group("/auth/order")
	orderApi.Use(util.JWT())
	{
		orderApi.POST("/buy", v1.Buy)
	}

	fundApi := r.Group("/auth/fund")
	fundApi.Use(util.JWT())
	{
		fundApi.POST("/charge", v1.Charge)
	}
	return r
}
