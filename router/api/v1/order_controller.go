package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"seckill-server/pkg/app"
	"seckill-server/service"
)

type BuyForm struct {
	FkGoodId   int64 `json:"fk_good_id"`
	Price      int64 `json:"price"`
	Num        int   `json:"num"`
	TotalPrice int64 `json:"total_price"`
	Address    struct {
		PkId     int    `json:"pk_id"`
		Province string `json:"province"`
		City     string `json:"city"`
		Detail   string `json:"detail"`
	} `json:"address"`
}

// @Tags Order
// @Summary 下单
// @Description 下单
// @Produce  json
// @Param form body BuyForm true "reqBody"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/order/buy [post]
// @Security Token
func Buy(c *gin.Context) {
	user := currentUser(c)
	if user == nil {
		return
	}
	var form BuyForm
	err := c.ShouldBindJSON(&form)
	if err != nil {
		app.ErrorResp(c, http.StatusBadRequest, err.Error())
		return
	}
	orderService := service.Order{
		FkGoodId:        form.FkGoodId,
		FkUserId:        user.PkId,
		Price:           form.Price,
		Num:             form.Num,
		UserName:        user.UserName,
		DeliveryAddress: fmt.Sprintf("%s%s%s", form.Address.Province, form.Address.City, form.Address.Detail),
	}
	order, err := orderService.Buy()
	if err != nil {
		app.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	app.SuccessResp(c, order)

}
