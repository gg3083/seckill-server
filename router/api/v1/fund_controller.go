package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seckill-server/pkg/app"
	"seckill-server/service"
)

type ChargeForm struct {
	Amount int `json:"amount"`
	Source int `json:"source"`
}

// @Tags Fund
// @Summary 充值
// @Description 充值
// @Produce  json
// @Param form body ChargeForm true "reqBody"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/fund/charge [post]
// @Security Token
func Charge(c *gin.Context) {
	user := currentUser(c)
	if user == nil {
		return
	}
	var form ChargeForm
	if err := c.ShouldBindJSON(&form); err != nil {
		app.ErrorResp(c, http.StatusBadRequest, "参数有误！")
		return
	}
	if form.Amount <= 0 {
		app.ErrorResp(c, http.StatusInternalServerError, "充值的金额不能小于0！")
		return
	}
	fundService := service.UserFund{
		FkUserId: user.PkId,
		Source:   form.Source,
	}
	amount := int64(form.Amount) * 10000
	err := fundService.AddBalance(amount)
	if err != nil {
		app.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	app.SuccessResp(c, nil)
}
