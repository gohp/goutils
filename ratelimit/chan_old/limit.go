package chan_old

import (
	"time"
)

type TokenBucket struct {
	fillInterval time.Duration
	capacity     int64
	Bucket       chan struct{}
}

func New(capacity int64, fillInterval time.Duration) *TokenBucket {
	tb := &TokenBucket{
		fillInterval: fillInterval,
		capacity:     capacity,
	}
	tb.Bucket = make(chan struct{}, tb.capacity)
	go tb.fillToken()
	return tb
}

func (t *TokenBucket) fillToken() {
	c := time.NewTicker(t.fillInterval)
	for {
		select {
		case <-c.C:
			select {
			case t.Bucket <- struct{}{}:
			default:
			}
			//fmt.Printf("token count %d in %v\n", len(t.Bucket), time.Now().UTC())
		}
	}
}

func (t *TokenBucket) TakeOne() bool {
	if len(t.Bucket) == 0 {
		return false
	}
	<-t.Bucket
	return true
}
