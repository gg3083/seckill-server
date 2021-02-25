package service

import (
	"log"
	"seckill-server/model"
	"seckill-server/pkg/consts"
	"seckill-server/pkg/util"
)

type Order struct {
	FkGoodId        string `json:"fk_good_id"`
	FkUserId        string `json:"fk_user_id"`
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
	hasStock, goods, err := goodsService.ValidStock()
	if err != nil || !hasStock {
		return nil, err
	}

	//增加商品销量
	if err := goodsService.Update(o.Num, goods.Version); err != nil {
		return nil, err
	}
	//检测用户余额
	fundService := UserFund{
		FkUserId: o.FkUserId,
	}
	o.Price = util.Multiply10000(o.Price)
	if _, err := fundService.ValidBalance(o.Price * int64(o.Num)); err != nil {
		return nil, err
	}
	//扣钱
	if err := fundService.AddBalance(-(o.Price * int64(o.Num)), consts.AmountSub); err != nil {
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
	log.Printf("用户%s 购买成功\n", o.UserName)
	return &orderDao, err

}
