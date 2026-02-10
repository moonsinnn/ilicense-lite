package util

import (
	"fmt"
	"time"
)

func ParseExpireAt(expireAt string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(layout, expireAt, time.UTC)
	if err != nil {
		return time.Time{}, fmt.Errorf("expire_at 格式错误: %w", err)
	}
	return t, nil
}
