// Package timer
/*
@Coding : utf-8
@time : 2022/8/7 22:55
@Author : yizhigopher
@Software : GoLand
*/
package timer

import (
	"container/list"
	"distributed/delayTask/timeWheel/bucket"
	"sync/atomic"
	"unsafe"
)

type Timer struct {
	Expiration int64
	Task       func()

	B       unsafe.Pointer
	Element *list.Element
}

func (t *Timer) GetBucket() *bucket.Bucket {
	return (*bucket.Bucket)(atomic.LoadPointer(&t.B))
}

func (t *Timer) SetBucket(b *bucket.Bucket) {
	atomic.StorePointer(&t.B, unsafe.Pointer(b))
}

func (t *Timer) Stop() bool {
	stopped := false
	for b := t.GetBucket(); b != nil; b = t.GetBucket() {
		stopped = b.Remove(t)
	}
	return stopped
}
