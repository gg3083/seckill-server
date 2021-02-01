package model

import "fmt"

type UserInfo struct {
	PkId     int64  `json:"pk_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	//...
	Token      string `json:"token"`
	CreateTime string `json:"create_time"`
}

type UserAddress struct {
	PkId     int64  `json:"pk_id"`
	FkUserId int64  `json:"fk_user_id"`
	Province string `json:"province"`
	City     string `json:"city"`
	Detail   string `json:"detail"`
	//...
	CreateTime string `json:"create_time"`
}

type UserFund struct {
	PkId     int64 `json:"pk_id"`
	FkUserId int64 `json:"fk_user_id"`
	Balance  int64 `json:"balance"`
	Version  int   `json:"version"`
	//...
	UpdateTime string `json:"update_time"`
}

type UserFundRecord struct {
	PkId     int64 `json:"pk_id"`
	FkUserId int64 `json:"fk_user_id"`
	Amount   int64 `json:"amount"`
	Type     int   `json:"type"`
	Source   int   `json:"source"`
	//...
	CreateTime string `json:"create_time"`
}

func (u *UserInfo) InsertUser() error {
	sql := fmt.Sprintf(
		"INSERT INTO `t_user_info`(`pk_id`, `user_name`, `password`, `token`) VALUES (?, ?, ?, ?)")
	stmt, err := mdb.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.PkId, u.UserName, u.Password, u.Token)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserInfo) GetUserById(id interface{}) (*UserInfo, error) {
	sql := fmt.Sprintf(
		"select * from t_user_info where pk_id = ?")
	row, err := mdb.Query(sql, id)
	if err != nil {
		return nil, err
	}
	var user UserInfo
	if row.Next() {
		err = row.Scan(&user.PkId, &user.UserName, &user.Password, &user.Token, &user.CreateTime)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("id:%v 不存在", id)
	}
	return &user, nil
}

func (u *UserInfo) GetUserByName(userName string) (*UserInfo, error) {
	sql := fmt.Sprintf(
		"select * from t_user_info where user_name = ?")
	row, err := mdb.Query(sql, userName)
	if err != nil {
		return nil, err
	}
	var user UserInfo
	if row.Next() {
		err = row.Scan(&user.PkId, &user.UserName, &user.Password, &user.Token, &user.CreateTime)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("用户:%v 不存在", userName)
	}
	return &user, nil
}

func (u *UserInfo) UpdateTokenById(token, id interface{}) error {
	sql := fmt.Sprintf(
		"update t_user_info set token = ? where pk_id = ?")
	stmt, err := mdb.Prepare(sql)
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(token, id); err != nil {
		return err
	}

	return nil
}

func (u *UserAddress) InsertUserAddress() error {
	sql := fmt.Sprintf(
		"INSERT INTO `t_user_address`(`pk_id`, `fk_user_id`, `province`, `city`,`detail`) VALUES (?, ?, ?, ?,?)")
	stmt, err := mdb.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.PkId, u.FkUserId, u.Province, u.City, u.Detail)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserAddress) GetUserAddressForLast(userId interface{}) (*UserAddress, error) {
	sql := fmt.Sprintf(
		"select * from t_user_address where fk_user_id = ? order by pk_id desc limit 1")
	row, err := mdb.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	var userAddress UserAddress
	if row.Next() {
		err = row.Scan(&userAddress.PkId, &userAddress.FkUserId, &userAddress.Province, &userAddress.City, &userAddress.Detail, &userAddress.CreateTime)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("用户:%v 不存在", userId)
	}
	return &userAddress, nil
}

func (u *UserFund) InsertUserFund() error {
	sql := fmt.Sprintf(
		"INSERT INTO `t_user_fund`(`pk_id`, `fk_user_id`, `balance`, `version`) VALUES (?, ?, ?, 0)")
	stmt, err := mdb.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.PkId, u.FkUserId, u.Balance)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserFund) GetUserFundByUserId(userId interface{}) (*UserFund, error) {
	sql := fmt.Sprintf(
		"select * from t_user_fund where fk_user_id = ?")
	row, err := mdb.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	var userFund UserFund
	if row.Next() {
		err = row.Scan(&userFund.PkId, &userFund.FkUserId, &userFund.Balance, &userFund.Version, &userFund.UpdateTime)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("用户:%v 不存在", userId)
	}
	return &userFund, nil
}

func (u *UserFund) AddFundBalanceByUserId(userId interface{}, amount int64) error {
	sql := fmt.Sprintf(
		"update t_user_fund set balance = balance + ? , version = version + 1 where fk_user_id = ?")
	stmt, err := mdb.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(amount, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserFundRecord) InsertUserFundRecord() error {
	sql := fmt.Sprintf(
		"INSERT INTO `t_user_fund_record` (`pk_id`, `fk_user_id`, `amount`, `type`,`source`) VALUES (?, ?, ?, ?, ?)")
	stmt, err := mdb.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.PkId, u.FkUserId, u.Amount, u.Type, u.Source)
	if err != nil {
		return err
	}
	return nil
}
