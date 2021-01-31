package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seckill-server/pkg/app"
	"seckill-server/service"
)

type RegisterForm struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

// @Tags User
// @Summary 注册
// @Description 注册
// @Produce  json
// @Param form body service.User true "reqBody"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	var from RegisterForm

	if err := c.ShouldBindJSON(&from); err != nil {
		app.ErrorResp(c, http.StatusBadRequest, "参数有误！")
		return
	}
	userService := service.User{
		UserName: from.UserName,
		PassWord: from.PassWord,
	}
	info, err := userService.Register()
	if err != nil {
		app.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	app.SuccessResp(c, info)
}

// @Tags User
// @Summary 登录
// @Description 登录
// @Produce  json
// @Param form body service.User true "reqBody"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var from RegisterForm

	if err := c.ShouldBindJSON(&from); err != nil {
		app.ErrorResp(c, http.StatusBadRequest, "参数有误！")
		return
	}
	userService := service.User{
		UserName: from.UserName,
		PassWord: from.PassWord,
	}
	info, err := userService.Login()
	if err != nil {
		app.ErrorResp(c, http.StatusBadRequest, err.Error())
		return
	}
	app.SuccessResp(c, info)
}
