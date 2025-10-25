package utils

import "time"

type Timer struct {
	startTime time.Time
	endTime   time.Time
	running   bool
}

func (t *Timer) Start() {
	t.startTime = time.Now()
	t.running = true
}

func (t *Timer) Stop() {
	if t.running {
		t.endTime = time.Now()
		t.running = false
	}
}

func (t *Timer) Elapsed() time.Duration {
	if t.running {
		return time.Since(t.startTime)
	}
	return t.endTime.Sub(t.startTime)
}

func (t *Timer) Reset() {
	t.startTime = time.Time{}
	t.endTime = time.Time{}
	t.running = false
}