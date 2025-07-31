package clock



import (
	"fmt"
)

type Clock struct {
	m int
	h int
}

func New(h, m int) Clock {

	if m >= 60 {
		h += m / 60
		m = m % 60
	}
	if h >= 24 {
		h = h % 24
	}

	for m < 0 {
		m = 60 + m
		h--
	}
	for h < 0 {
		h = 24 + h
	}
	return Clock{m: m, h: h}
}

func (c Clock) Add(m int) Clock {
	m = c.m + m
	h := c.h
	if m >= 60 {
		h += m / 60
		m = m % 60
	}
	if h >= 24 {
		h = h % 24
	}
	return Clock{m: m, h: h}
}

func (c Clock) Subtract(m int) Clock {
	minute := c.m - m
	h := c.h
	for minute < 0 {
		minute = 60 + minute
		h--
	}
    for h < 0 {
		h = 24 + h
	}
	return Clock{m: minute, h: h}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
