package service

import (
	"errors"
	"seckill-server/model"
	"seckill-server/pkg/util"
)

type Goods struct {
	PkId        string
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

func (g *Goods) ValidStock() (bool, *model.Goods, error) {
	goods, err := g.Get()
	if err != nil {
		return false, nil, err
	}
	if goods.Stock <= 0 {
		return false, nil, errors.New("库存不足！")
	}
	return true, goods, nil
}

func (g *Goods) Add() error {
	g.Price = util.Multiply10000(g.Price)
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

func (g *Goods) Update(saleNum, version int) error {
	goodsModel := model.Goods{
		PkId:    g.PkId,
		Version: version,
	}
	return goodsModel.UpdateSaleNum(saleNum, version)
}
