// Package delayQueue
/*
@Coding : utf-8
@time : 2022/8/7 22:36
@Author : yizhigopher
@Software : GoLand
*/
package delayQueue

import (
	"container/heap"
	"distributed/delayTask/timeWheel/priorityQueue"
	"sync"
	"sync/atomic"
	"time"
)

type DelayQueue struct {
	c chan interface{}

	mu sync.Mutex
	pq priorityQueue.PriorityQueue

	sleeping int32
	wakeupC  chan struct{}
}

func New(size int) *DelayQueue {
	return &DelayQueue{
		c:       make(chan interface{}),
		pq:      priorityQueue.NewPriorityQueue(size),
		wakeupC: make(chan struct{}),
	}
}

func (dq *DelayQueue) Offer(elem interface{}, expiration int64) {
	item := &priorityQueue.Item{
		Value:    elem,
		Priority: expiration,
	}

	dq.mu.Lock()
	heap.Push(&dq.pq, item)
	index := item.Index
	dq.mu.Unlock()

	if index == 0 {
		if atomic.CompareAndSwapInt32(&dq.sleeping, 1, 0) {
			dq.wakeupC <- struct{}{}
		}
	}
}

func (dq *DelayQueue) Poll(exitC chan struct{}, nowF func() int64) {
	for {
		now := nowF()
		dq.mu.Lock()
		item, delta := dq.pq.PeekAndShift(now)

		if item == nil {
			atomic.StoreInt32(&dq.sleeping, 1)
		}
		dq.mu.Unlock()
		if item == nil {
			if delta == 0 {
				select {
				case <-dq.wakeupC:
					continue
				case <-exitC:
					goto exit
				}
			} else if delta > 0 {
				select {
				case <-dq.wakeupC:
					continue
				case <-time.After(time.Duration(delta) * time.Millisecond):
					if atomic.SwapInt32(&dq.sleeping, 0) == 0 {
						<-dq.wakeupC
					}
					continue
				case <-exitC:
					goto exit
				}
			}
		}
		select {
		case dq.c <- item.Value:
		case <-exitC:
			goto exit
		}
	}
exit:
	atomic.StoreInt32(&dq.sleeping, 0)
}
