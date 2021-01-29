package model

type UserInfo struct {
	PkUserId int    `json:"pk_user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	//...
	Token      string `json:"token"`
	CreateTime string `json:"create_time"`
}

type UserAddress struct {
	PkId     int    `json:"pk_id"`
	FkUserId int    `json:"fk_user_id"`
	Province string `json:"province"`
	City     string `json:"city"`
	Detail   string `json:"detail"`
	//...
	CreateTime string `json:"create_time"`
}

type UserFund struct {
	PkId     int   `json:"pk_id"`
	FkUserId int   `json:"fk_user_id"`
	Balance  int64 `json:"balance"`
	Version  int   `json:"version"`
	//...
	UpdateTime string `json:"update_time"`
}

type UserFundRecord struct {
	PkId         int   `json:"pk_id"`
	FkUserFundId int   `json:"fk_user_fund_id"`
	Amount       int64 `json:"amount"`
	Type         int   `json:"type"`
	//...
	CreateTime string `json:"create_time"`
}
