package service

import (
	"seckill-server/model"
	"seckill-server/pkg/consts"
	"seckill-server/pkg/util"
)

type Order struct {
	FkGoodId        int64  `json:"fk_good_id"`
	FkUserId        int64  `json:"fk_user_id"`
	GoodsName       string `json:"goods_name"`
	Price           int64  `json:"price"`
	Num             int    `json:"num"`
	TotalPrice      int64  `json:"total_price"`
	UserName        string `json:"user_name"`
	DeliveryAddress string `json:"delivery_address"`
	PayTime         int64  `json:"pay_time"`
}

func (o *Order) Buy() {
	//检查库存
	goodsService := Goods{
		PkId: o.FkGoodId,
	}
	goods, err := goodsService.Get()
	if err != nil {
		return
	}
	if goods.Stock <= 0 {
		return
	}
	//检测用户余额
	fundDao := model.UserFund{}
	fund, err := fundDao.GetUserFundByUserId(o.FkUserId)
	if err != nil {
		return
	}
	if fund.Balance < o.Price*int64(o.Num) {
		return
	}
	id := util.GetUniqueNo(consts.BusinessOrderTable)
	orderDao := model.Order{
		PkId:            id,
		FkGoodId:        o.FkGoodId,
		FkUserId:        o.FkUserId,
		GoodsName:       goods.GoodsName,
		Price:           goods.Price,
		Num:             o.Num,
		TotalPrice:      goods.Price * int64(o.Num),
		UserName:        "userName",
		DeliveryAddress: o.DeliveryAddress,
	}

	if err := orderDao.Insert(); err != nil {
		return
	}
	//生成订单
}
