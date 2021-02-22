package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"seckill-server/pkg/app"
	"seckill-server/pkg/consts"
	"seckill-server/pkg/util"
	"seckill-server/request"
	"seckill-server/service"
	"strconv"
	"time"
)

// @Tags Project
// @Summary 查询
// @Description 列出全部项目
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json model.Goods
// @Router /goods/{id} [get]
func GetGoods(c *gin.Context) {
	paramId := c.Param("id")
	if paramId == "" {
		app.ErrorResp(c, http.StatusBadRequest, "Id不能为空")
		return
	}
	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		app.ErrorResp(c, http.StatusBadRequest, "Id错误")
		return
	}
	goodsService := service.Goods{
		PkId: id,
	}
	res, err := goodsService.Get()
	if err != nil {
		app.ErrorResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	app.SuccessResp(c, res)
}

// @Tags Project
// @Summary 添加
// @Description 添加
// @Produce  json
// @Param form body request.Goods true "reqBody"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /goods/ [post]
func AddGoods(c *gin.Context) {
	var from request.Goods

	if err := c.ShouldBindJSON(&from); err != nil {
		app.ErrorResp(c, http.StatusBadRequest, "请求格式错误！")
		return
	}
	seckillTime, err := time.Parse("2006-01-02 15:04:05", from.SeckillTime)
	if err != nil {
		app.ErrorResp(c, http.StatusBadRequest, "秒杀开始时间格式错误！")
		return
	}
	id := util.GetUniqueNo(consts.BusinessGoodsTable)
	goodsService := service.Goods{
		PkId:        id,
		GoodsName:   from.GoodsName,
		Price:       from.Price,
		Stock:       from.Stock,
		IsSeckill:   1,
		SeckillTime: seckillTime.Unix(),
	}
	if err := goodsService.Add(); err != nil {
		app.ErrorResp(c, http.StatusBadRequest, err.Error())
		return
	}
	app.SuccessResp(c, fmt.Sprintf("%v", id))
}
