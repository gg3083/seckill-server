package util

import (
	"crypto/md5"
	"encoding/hex"
	"seckill-server/pkg/consts"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value, salt string) string {
	m := md5.New()
	m.Write([]byte(value))
	m.Write([]byte(salt))
	return hex.EncodeToString(m.Sum(nil))
}

// 判断str的MD5加密是否为md5Str
func MD5Equals(str, salt, md5Str string) bool {
	return EncodeMD5(str, salt) == md5Str
}

//根据用户名、盐 在加一次密得到一个盐
func GetMd5String(str string) string {
	salt := consts.PasswordSalt
	return EncodeMD5(str, salt)
}
