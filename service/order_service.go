package service

import (
	"seckill-server/model"
	"seckill-server/pkg/consts"
	"seckill-server/pkg/util"
)

type Order struct {
	FkGoodId        int64  `json:"fk_good_id"`
	FkUserId        int64  `json:"fk_user_id"`
	Price           int64  `json:"price"`
	Num             int    `json:"num"`
	TotalPrice      int64  `json:"total_price"`
	UserName        string `json:"user_name"`
	DeliveryAddress string `json:"delivery_address"`
	PayTime         int64  `json:"pay_time"`
}

func (o *Order) Buy() (*model.Order, error) {
	//检查库存
	goodsService := Goods{
		PkId: o.FkGoodId,
	}
	if _, err := goodsService.ValidStock(); err != nil {
		return nil, err
	}

	goods, err := goodsService.Get()

	if err != nil {
		return nil, err
	}

	//检测用户余额
	fundService := UserFund{
		FkUserId: o.FkUserId,
	}
	if _, err := fundService.ValidBalance(o.Price * int64(o.Num)); err != nil {
		return nil, err
	}
	//扣钱
	if err := fundService.AddBalance(-(o.Price * int64(o.Num)), consts.AmountSub); err != nil {
		return nil, err
	}
	//增加商品销量
	if err := goodsService.Update(o.Num); err != nil {
		return nil, err
	}

	//生成订单
	id := util.GetUniqueNo(consts.BusinessOrderTable)
	orderDao := model.Order{
		PkId:            id,
		FkGoodId:        o.FkGoodId,
		FkUserId:        o.FkUserId,
		GoodsName:       goods.GoodsName,
		Price:           goods.Price,
		Num:             o.Num,
		TotalPrice:      goods.Price * int64(o.Num),
		UserName:        o.UserName,
		DeliveryAddress: o.DeliveryAddress,
	}

	if err := orderDao.Insert(); err != nil {
		return nil, err
	}
	return &orderDao, err

}
