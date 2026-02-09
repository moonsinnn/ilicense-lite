/*
 * Date: 2022/10/24
 * File: queue.go
 */

// Package lmq TODO package function desc
package lmq

import (
	"fmt"
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

// LocalMessageQueue
// @Description: 简单的本地消息队列，主要用于错误重试
type LocalMessageQueue struct {
	size     int                // 队列大型
	queue    chan *LocalMessage // 队列缓存
	state    int32              // 消费协程运行状态，0停止，1运行中
	idleTime time.Duration      // 消费协程空转时间
	pushTime time.Time          // 添加消息时间
}

func NewLocalMessageQueue() *LocalMessageQueue {
	return &LocalMessageQueue{
		idleTime: defaultIdleTime,
		size:     defaultQueueSize,
		queue:    make(chan *LocalMessage, defaultQueueSize),
	}
}

// Push
// @Description: 添加消息
// @receiver lmq
func (lmq *LocalMessageQueue) Push(message *LocalMessage) {
	if lmq.full() {
		return
	}
	if !lmq.running() {
		lmq.run()
	}
	lmq.queue <- message
	lmq.pushTime = time.Now()
}

// close
// @Description: 关闭消费协程
// @receiver lmq
func (lmq *LocalMessageQueue) close() {
	atomic.AddInt32(&lmq.state, -1)
}

// closed
// @Description: 消费协程是否已关闭
// @receiver lmq
// @return bool
func (lmq *LocalMessageQueue) closed() bool {
	return atomic.LoadInt32(&lmq.state) == 0
}

// empty
// @Description: 消息队列是否为空
// @receiver lmq
// @return bool
func (lmq *LocalMessageQueue) empty() bool {
	return 0 == len(lmq.queue)
}

// full
// @Description: 消息队列是否已满
// @receiver lmq
// @return bool
func (lmq *LocalMessageQueue) full() bool {
	return lmq.size == len(lmq.queue)
}

// run
// @Description: 运行消费协程
// @receiver lmq
func (lmq *LocalMessageQueue) run() {
	atomic.AddInt32(&lmq.state, 1)
	go lmq.start()
}

// running
// @Description: 消费协程是否运行中
// @receiver lmq
// @return bool
func (lmq *LocalMessageQueue) running() bool {
	return atomic.LoadInt32(&lmq.state) == 1
}

// expired
// @Description: 消费协程是否已空闲指定时间
// @receiver lmq
// @return bool
func (lmq *LocalMessageQueue) expired() bool {
	return time.Now().After(lmq.pushTime.Add(lmq.idleTime))
}

// start
// @Description: 启动消费协程
// @receiver lmq
func (lmq *LocalMessageQueue) start() {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			stack := make([]byte, 2<<10)
			length := runtime.Stack(stack, false)
			realErr := fmt.Sprintf("%v %s\n", err, stack[:length])
			log.Fatalf("lmq is panic: %v", realErr)
		}
	}()
	ticker := time.NewTicker(defaultCheckTime)
	defer ticker.Stop()
	for {
		select {
		case message := <-lmq.queue:
			if message.count < message.retry {
				message.count++
				if ok := message.process(); !ok {
					lmq.Push(message)
				}
			} else {
				log.Fatalf("lmq message is exceed max time, message_id:%+v", message.GetID())
			}
		case <-ticker.C:
			if lmq.empty() && lmq.expired() {
				lmq.close()
				return
			}
		}
	}
}
