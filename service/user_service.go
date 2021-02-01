package service

import (
	"github.com/pkg/errors"
	"seckill-server/model"
	"seckill-server/pkg/consts"
	"seckill-server/pkg/util"
)

type User struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func (user *User) Register() (*model.UserInfo, error) {

	salt := util.GetMd5String(user.UserName)
	passwordSalt := util.EncodeMD5(user.PassWord, salt)
	userId := util.GetUniqueNo(consts.BusinessUserTable)
	userDao := model.UserInfo{
		PkId:     userId,
		UserName: user.UserName,
		Password: passwordSalt,
	}
	info, _ := userDao.GetUserByName(user.UserName)
	if info != nil {
		return nil, errors.New("用户名已存在!")
	}
	//生成token
	token, err := util.GenerateToken(userId, user.UserName)
	if err != nil {
		return nil, errors.New("创建token失败!")
	}
	userDao.Token = token
	if err := userDao.InsertUser(); err != nil {
		return nil, err
	}
	// 初始化余额
	fundId := util.GetUniqueNo(consts.BusinessUserFundTable)
	fundDao := model.UserFund{
		PkId:     fundId,
		FkUserId: userId,
		Balance:  0,
	}
	if err := fundDao.InsertUserFund(); err != nil {
		return nil, err
	}
	return &userDao, nil
}

func (user *User) Login() (*model.UserInfo, error) {

	salt := util.GetMd5String(user.UserName)
	userDao := model.UserInfo{}
	info, _ := userDao.GetUserByName(user.UserName)
	if info == nil {
		return nil, errors.New("用户不存在!")
	}
	if !util.MD5Equals(user.PassWord, salt, info.Password) {
		return nil, errors.New("密码不正确!")
	}
	//重新生成token
	token, err := util.GenerateToken(info.PkId, info.UserName)
	if err != nil {
		return nil, errors.New("创建token失败!")
	}
	info.Token = token
	if err := userDao.UpdateTokenById(token, info.PkId); err != nil {
		return nil, err
	}
	return info, nil
}

func (user *User) GetByPkId(id interface{}) (*model.UserInfo, error) {

	userDao := model.UserInfo{}
	info, _ := userDao.GetUserById(id)
	if info == nil {
		return nil, errors.New("用户不存在!")
	}
	return info, nil
}

type UserAddress struct {
	FkUserId int64
	Province string
	City     string
	Detail   string
}

func (u *UserAddress) InsertUserAddress() (int64, error) {
	id := util.GetUniqueNo(consts.BusinessUserAddressTable)
	addressDao := model.UserAddress{
		PkId:     id,
		FkUserId: u.FkUserId,
		Province: u.Province,
		City:     u.City,
		Detail:   u.Detail,
	}

	if err := addressDao.InsertUserAddress(); err != nil {
		return id, err
	}
	return id, nil
}
