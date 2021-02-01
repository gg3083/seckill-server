package util

import (
	"seckill-server/pkg/consts"
	"time"
)

func GetUniqueNo(i int) int64 {
	switch i {
	case consts.BusinessUserTable:
		return time.Now().UnixNano()
	case consts.BusinessGoodsTable:
		return time.Now().UnixNano()
	case consts.BusinessUserFundRecordTable:
		return time.Now().UnixNano()
	case consts.BusinessUserFundTable:
		return time.Now().UnixNano()
	case consts.BusinessUserAddressTable:
		return time.Now().UnixNano()
	case consts.BusinessOrderTable:
		return time.Now().UnixNano()
	default:
		return time.Now().UnixNano()
	}
}
