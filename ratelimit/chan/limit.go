package _chan

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	fillInterval time.Duration
	capacity     int64
	Bucket       chan struct{}
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
			fmt.Printf("token count %d in %v\n", len(t.Bucket), time.Now().UTC())
		}
	}
}

func main() {
	done := make(chan struct{})
	tb := &TokenBucket{
		fillInterval: time.Millisecond * 10,
		capacity:     100,
	}
	tb.Bucket = make(chan struct{}, tb.capacity)

	go tb.fillToken()
	<-done
}
