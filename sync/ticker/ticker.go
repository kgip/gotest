package ticker

import (
	"time"
)

type Ticker struct {
	ticker *time.Ticker
	callback func()
	stop bool
}

func NewTicker(duration time.Duration,callback func()) *Ticker {
	return &Ticker{
		ticker: time.NewTicker(duration),
		callback: callback,
		stop: false,
	}
}

//执行定时器
func (ticker *Ticker) Do() *Ticker {
	go func() {
		for !ticker.stop {
			<-ticker.ticker.C
			ticker.callback()
		}
	}()
	return ticker
}

/**
    停止执行定时器
 */
func (ticker *Ticker) Stop()  {
	ticker.stop = true
	ticker.ticker.Stop()
}
