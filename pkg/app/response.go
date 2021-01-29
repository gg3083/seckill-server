package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// ErrorResp 错误返回值
func ErrorResp(c *gin.Context, code int, errMsg string) {
	resp(c, http.StatusOK, code, errMsg, nil)
}

// ErrorResp 错误返回值
func UnauthorizedResp(c *gin.Context, code int, errMsg string) {
	resp(c, http.StatusUnauthorized, code, errMsg, nil)
}

// SuccessResp 正确返回值
func SuccessRespByCode(c *gin.Context, code int, data interface{}) {
	resp(c, http.StatusOK, code, "", data)
}

// SuccessResp 正确返回值
func SuccessResp(c *gin.Context, data interface{}) {
	resp(c, http.StatusOK, 0, "", data)
}

// resp 返回
func resp(c *gin.Context, httpCode, code int, msg string, data interface{}) {
	resp := Response{
		Code:   code,
		Msg:    msg,
		Data:   data,
	}
	c.Set("log_response:", &resp)
	c.JSON(httpCode, resp)
}
