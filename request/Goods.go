package request

type Goods struct {
	GoodsName   string `json:"goods_name"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
	SeckillTime string  `json:"seckill_time"`
}
