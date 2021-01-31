package model

type Order struct {
	PkId            int64  `json:"pk_id"`
	FkGoodId        int64  `json:"fk_good_id"`
	FkUserId        int64  `json:"fk_user_id"`
	GoodsName       string `json:"goods_name"`
	Price           int64  `json:"price"`
	Num             int    `json:"num"`
	TotalPrice      int64  `json:"total_price"`
	UserName        string `json:"user_name"`
	DeliveryAddress string `json:"delivery_address"`
	PayTime         int64  `json:"pay_time"`
	CreateTime      string `json:"create_time"`
}

func (order *Order) Insert() error {

}
