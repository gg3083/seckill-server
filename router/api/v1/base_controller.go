package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seckill-server/model"
	"seckill-server/pkg/app"
	"seckill-server/pkg/consts"
	"seckill-server/pkg/util"
	"seckill-server/service"
)

func currentUser(c *gin.Context) *model.UserInfo {
	claims, err := util.ParseToken(c.GetHeader(consts.HeaderToken))
	if err != nil {
		app.ErrorResp(c, http.StatusBadRequest, err.Error())
		return nil
	}
	userServer := service.User{}
	user, err := userServer.GetByPkId(claims.PkId)
	if err != nil {
		app.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return nil
	}
	return user
}
