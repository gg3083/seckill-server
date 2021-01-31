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
	id := util.GetUniqueNo(consts.BusinessUserTable)
	userDao := model.UserInfo{
		PkId:     id,
		UserName: user.UserName,
		Password: passwordSalt,
	}
	info, _ := userDao.GetUserByName(user.UserName)
	if info != nil {
		return nil, errors.New("用户名已存在!")
	}
	//生成token
	token, err := util.GenerateToken(info.PkId, info.UserName)
	if err != nil {
		return nil, errors.New("创建token失败!")
	}
	userDao.Token = token
	if err := userDao.InsertUser(); err != nil {
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
