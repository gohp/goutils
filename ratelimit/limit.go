package ratelimit

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity     int64
	fillInterval time.Duration
	mu           sync.Mutex

	q int64     // Quantity of token each time
	t time.Time // last operation time
	k int64     // last operation quantity of token
}

func New(capacity int64, fillInterval time.Duration) *TokenBucket {
	tb := &TokenBucket{
		capacity:     capacity,
		fillInterval: fillInterval,
		q:            1,
		t:            time.Now(),
		k:            0,
	}
	return tb
}

func (tb *TokenBucket) Take(count int64) bool {
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
		tb.k = tb.k - count
		return true
	}

	return false
}
