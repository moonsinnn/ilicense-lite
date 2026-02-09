/*
 * Date: 2022/10/24
 * File: base.go
 */

// Package lmq TODO package function desc
package lmq

import "time"

const (
	defaultRetryCount = 3                // 默认重试次数
	defaultQueueSize  = 100              // 默认队列大小
	defaultIdleTime   = 10 * time.Minute // 默认空闲时间
	defaultCheckTime  = 1 * time.Minute  // 默认检查时间频率
)

var (
	// 默认本地消息队列
	DefaultLocalMessageQueue = NewLocalMessageQueue()
)
