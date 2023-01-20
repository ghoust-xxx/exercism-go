package clock

import (
	"fmt"
)

type Clock struct {
	hours, minutes int
}

func New(h, m int) Clock {
	time := Clock {
		hours: h,
		minutes: m,
	}
	time.hours += time.minutes / 60
	time.minutes %= 60
	if time.minutes < 0 {
		time.minutes += 60
		time.hours--
	}
	time.hours %= 24
	if time.hours < 0 {
		time.hours += 24
	}
	return time
}

func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes + m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
