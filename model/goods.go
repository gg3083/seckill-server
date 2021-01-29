package model

import (
	"fmt"
)

type Goods struct {
	PkId        int64  `json:"pk_id"`
	GoodsName   string `json:"goods_name"`
	Price       int64  `json:"price"`
	SaleNum     int    `json:"sale_num"`
	Stock       int    `json:"stock"`
	IsSeckill   int    `json:"is_seckill"`
	SeckillTime int64  `json:"seckill_time"`
	Version     int    `json:"version"`
	//...
	UpdateTime string `json:"update_time"`
}

func (g *Goods) Insert() (int64, error) {
	sql := fmt.Sprintf(
		"INSERT INTO `t_goods`(`pk_id`, `goods_name`, `price`, `sale_num`, `stock`, `is_seckill`, `seckill_time`, `version`) VALUES (?, ?, ?, 0, ?, ?, ?, 0)")
	stmt, err := mdb.Prepare(sql)
	if err != nil {
		return -1, err
	}
	_, err = stmt.Exec(g.PkId, g.GoodsName, g.Price,  g.Stock, g.IsSeckill, g.SeckillTime)
	if err != nil {
		return -1, err
	}
	return g.PkId, nil
}

func (g *Goods) GetByPkId(id interface{}) (*Goods, error ){
	sql := fmt.Sprintf(
		"select * from t_goods where pk_id = ?")
	row, err := mdb.Query(sql, id)
	if err != nil {
		return nil, err
	}
	var goods Goods
	if row.Next() {
		err = row.Scan(&goods.PkId, &goods.GoodsName, &goods.Price, &goods.SaleNum, &goods.Stock, &goods.IsSeckill, &goods.SeckillTime, &goods.Version, &goods.UpdateTime)
		if err != nil {
			return nil, err
		}
	}else {
		return nil, fmt.Errorf("id:%v 不存在", id)
	}
	return &goods, nil
}
