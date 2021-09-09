package timer

import (
	"log"
	"time"
)

type Timer struct {
	timer *time.Timer
	callback func()
}

//创建一个定时器
func NewTimer(duration time.Duration,callback func()) *Timer {
	return &Timer{
		timer: time.NewTimer(duration),
		callback: callback,
	}
}

func (this *Timer) Callback(callback func())  {
	this.callback = callback
}

func (this *Timer) Do() {
	msg := <-this.timer.C
	this.callback()
	log.Println(msg)
}