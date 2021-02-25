package model

import "fmt"

type Order struct {
	PkId            string `json:"pk_id"`
	FkGoodId        string `json:"fk_good_id"`
	FkUserId        string `json:"fk_user_id"`
	GoodsName       string `json:"goods_name"`
	Price           int64  `json:"price"`
	Num             int    `json:"num"`
	TotalPrice      int64  `json:"total_price"`
	UserName        string `json:"user_name"`
	DeliveryAddress string `json:"delivery_address"`
	PayTime         int64  `json:"pay_time"`
	CreateTime      string `json:"create_time"`
}

func (o *Order) Insert() error {
	sql := fmt.Sprintf(
		"INSERT INTO `t_order`(`pk_id`, `fk_good_id`, `fk_user_id`, `goods_name`, `price`, `num`, `total_price`, `user_name`,`delivery_address`,`pay_time`) VALUES (?, ?, ?, ?, ?, ?, ?,?, ?, 0)")
	stmt, err := mdb.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(o.PkId, o.FkGoodId, o.FkUserId, o.GoodsName, o.Price, o.Num, o.TotalPrice, o.UserName, o.DeliveryAddress)
	if err != nil {
		return err
	}
	return nil
}
