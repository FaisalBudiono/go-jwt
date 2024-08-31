package nower

import "time"

type Nower interface {
	Now() time.Time
}

func New() *nowerReal {
	return &nowerReal{}
}

func NewFake(t time.Time) *nowerFake {
	return &nowerFake{
		now: t,
	}
}

type nowerReal struct{}

func (n *nowerReal) Now() time.Time {
	return time.Now()
}

type nowerFake struct {
	now time.Time
}

func (n *nowerFake) Now() time.Time {
	return n.now
}
