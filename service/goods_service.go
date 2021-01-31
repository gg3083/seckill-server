package service

import (
	"seckill-server/model"
)

type Goods struct {
	PkId        interface{}
	GoodsName   string
	Price       int64
	Stock       int
	IsSeckill   int
	SeckillTime int64
}

func (g *Goods) Get() (*model.Goods, error) {
	goodsModel := model.Goods{}
	return goodsModel.GetByPkId(g.PkId)
}

func (g *Goods) Add() error {
	goodsModel := model.Goods{
		PkId:        g.PkId.(int64),
		GoodsName:   g.GoodsName,
		Price:       g.Price,
		IsSeckill:   g.IsSeckill,
		SeckillTime: g.SeckillTime,
	}
	return goodsModel.Insert()
}
