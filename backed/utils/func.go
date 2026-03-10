package utils

import (
	"errors"
	"math"
)

// 安全地将 uint64 转换为 int64
func Uint64ToInt64Safe(u uint64) (int64, error) {
	if u > math.MaxInt64 {
		return 0, errors.New("uint64 值超出 int64 范围")
	}
	return int64(u), nil
}
