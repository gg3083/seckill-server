package service

import (
	"errors"
	"seckill-server/model"
)

type Goods struct {
	PkId        int64
	GoodsName   string
	Price       int64
	Stock       int
	IsSeckill   int
	SeckillTime int64
	SaleNum     int
}

func (g *Goods) Get() (*model.Goods, error) {
	goodsModel := model.Goods{}
	return goodsModel.GetByPkId(g.PkId)
}

func (g *Goods) ValidStock() (bool, error) {
	goods, err := g.Get()
	if err != nil {
		return false, err
	}
	if goods.Stock <= 0 {
		return false, errors.New("库存不足！")
	}
	return true, nil
}

func (g *Goods) Add() error {
	goodsModel := model.Goods{
		PkId:        g.PkId,
		GoodsName:   g.GoodsName,
		Price:       g.Price,
		IsSeckill:   g.IsSeckill,
		SeckillTime: g.SeckillTime,
		Stock:       g.Stock,
	}
	return goodsModel.Insert()
}

func (g *Goods) Update(saleNum int) error {
	goodsModel := model.Goods{
		PkId: g.PkId,
	}
	return goodsModel.UpdateSaleNum(saleNum)
}
