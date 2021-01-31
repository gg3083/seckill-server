package controller

import "github.com/gin-gonic/gin"

type BuyForm struct {
	FkGoodId   int    `json:"fk_good_id"`
	FkUserId   int    `json:"fk_user_id"`
	GoodsName  string `json:"goods_name"`
	Price      int64  `json:"price"`
	Num        int    `json:"num"`
	TotalPrice int64  `json:"total_price"`
	UserName   string `json:"user_name"`
	Address    struct {
		PkId     int    `json:"pk_id"`
		Province string `json:"province"`
		City     string `json:"city"`
		Detail   string `json:"detail"`
	} `json:"address"`
}

func Buy(c *gin.Context) {

}
