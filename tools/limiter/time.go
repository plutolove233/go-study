package main

import (
	"fmt"
	"sync"
	"time"
)

type Limiter struct {
	mu        sync.Mutex
	rate      int
	timestamp time.Time
	interval  time.Duration
	requests  int
}

func newLimiter(rate int, interval time.Duration) *Limiter {
	return &Limiter{
		rate: rate,
		interval: interval,
		timestamp: time.Now(),
		requests: 0,
	}
}

func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if time.Since(l.timestamp) > l.interval {
		l.timestamp = time.Now()
		l.requests = 0
	}

	if l.requests < l.rate {
		l.requests += 1
		return true
	}

	return false
}

func main() {
	limiter := newLimiter(1000, 60*time.Second)
	for v := range 1100 {
		if limiter.Allow() {
			fmt.Println("request passed! ", v+1)
		} else {
			fmt.Println("request denied! ", v+1)
		}
	}
}
