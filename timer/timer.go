package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	H    string
	M    string
	S    string
	IsAM bool
	stop chan struct{}
	Tick chan struct{}
}

func New() *Timer {
	return &Timer{stop: make(chan struct{}), Tick: make(chan struct{})}
}

func (c *Timer) Start() {
	ticker := time.NewTicker(time.Second)
	t := <-ticker.C
	c.SetTime(t)
	go c.Watch(ticker)
}

func (c *Timer) Stop() {
	c.stop <- struct{}{}
}

func (c *Timer) SetTime(t time.Time) {
	h := uint8(t.Hour())
	if h >= 12 {
		h = h - 12
		c.IsAM = false
	} else {
		c.IsAM = true
	}
	c.H = fmt.Sprintf("%02d", h)
	c.M = fmt.Sprintf("%02d", uint8(t.Minute()))
	c.S = fmt.Sprintf("%02d", uint8(t.Second()))
	c.Tick <- struct{}{}
}

func (c *Timer) Watch(ticker *time.Ticker) {
	for {
		select {
		case t := <-ticker.C:
			c.SetTime(t)
		case <-c.stop:
			ticker.Stop()
			return
		}
	}
}
