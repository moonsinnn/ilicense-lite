/*
 * Date: 2022/10/24
 * File: message.go
 */

// Package lmq TODO package function desc
// @ignore
package lmq

// LocalMessage
// @Description: 本地消息
type LocalMessage struct {
	id      string      // 消息ID或会话ID
	count   int         // 当前执行次数
	retry   int         // 最大重试次数
	process func() bool // 执行处理函数
}

// NewLocalMessage
// @Description: 创建消息
// @return *LocalMessage
// @ignore
func NewLocalMessage(id string, process func() bool) *LocalMessage {
	return &LocalMessage{id: id, retry: defaultRetryCount, process: process}
}

// GetID ID
// @Description: 消息ID或会话ID
// @receiver lmq
// @return string
// @ignore
func (lmq *LocalMessage) GetID() string {
	return lmq.id
}
