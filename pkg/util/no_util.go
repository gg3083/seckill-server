package util

import (
	"fmt"
	"math/rand"
	"seckill-server/pkg/consts"
	"time"
)

func GetUniqueNo(i int) string {
	var timeNo int64 = 0
	switch i {
	case consts.BusinessUserTable:
		timeNo = time.Now().UnixNano()
	case consts.BusinessGoodsTable:
		timeNo = time.Now().UnixNano()
	case consts.BusinessUserFundRecordTable:
		timeNo = time.Now().UnixNano()
	case consts.BusinessUserFundTable:
		timeNo = time.Now().UnixNano()
	case consts.BusinessUserAddressTable:
		timeNo = time.Now().UnixNano()
	case consts.BusinessOrderTable:
		timeNo = time.Now().UnixNano()
	default:
		timeNo = time.Now().UnixNano()
	}
	return fmt.Sprintf("%v%v", timeNo, rand.Intn(1999)-1000)
}
