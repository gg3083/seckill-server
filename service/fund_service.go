package service

import (
	"errors"
	"seckill-server/model"
	"seckill-server/pkg/consts"
	"seckill-server/pkg/util"
)

type UserFund struct {
	FkUserId int64
	Source   int
}

func (fund *UserFund) ValidBalance(amount int64) (bool, error) {

	//检测用户余额
	fundDao := model.UserFund{}
	userFund, err := fundDao.GetUserFundByUserId(fund.FkUserId)
	if err != nil {
		return false, err
	}
	if userFund.Balance < amount {
		return false, errors.New("余额不足！")
	}

	return true, nil
}

func (fund *UserFund) AddBalance(amount int64) error {

	fundDao := model.UserFund{}

	if err := fundDao.AddFundBalanceByUserId(fund.FkUserId, amount); err != nil {
		return err
	}

	fundRecordId := util.GetUniqueNo(consts.BusinessUserFundRecordTable)
	recordDao := model.UserFundRecord{
		PkId:     fundRecordId,
		FkUserId: fund.FkUserId,
		Amount:   amount,
		Type:     consts.AmountAdd,
		Source:   fund.Source,
	}
	if err := recordDao.InsertUserFundRecord(); err != nil {
		return err
	}
	return nil
}
