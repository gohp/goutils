package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity     int64
	fillInterval time.Duration
	mu           sync.Mutex

	q int64
	t time.Time
	k int64
}

func (tb *TokenBucket) take(count int64) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	if count <= 0 {
		return true
	}
	curTime := time.Now()

	// (t2 - t1) / ti * x + k1
	tb.k += int64(curTime.Sub(tb.t)) * tb.q
	tb.t = curTime

	if tb.k > tb.capacity {
		tb.k = tb.capacity
	}

	if tb.k-count >= 0 {
		tb.k = tb.k-count
		return true
	}

	return false
}

func main() {
	var i = 1
	var fillInterval = time.Millisecond * 10

	bucket := &TokenBucket{
		capacity:     int64(i),
		fillInterval: fillInterval,
		q:            1,
		t:            time.Now(),
		k:            0,
	}

	takeResult := bucket.take(1)
	fmt.Println("take result", takeResult)
}
