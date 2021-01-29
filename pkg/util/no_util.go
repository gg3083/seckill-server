package util

import "time"

func GetUniqueNo(i int) int64 {
	switch i {
	case 1:
		return time.Now().UnixNano()
	case 2:
		return time.Now().UnixNano()
	default:
		return time.Now().UnixNano()
	}
}
