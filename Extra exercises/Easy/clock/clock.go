// Package clock implements a clock that handles times without dates.
package clock

import "strconv"

// Clock is a type of time.
type Clock struct {
	hour   int
	minute int
}

// New builds a time.
func New(hour, minute int) (c Clock) {
	for minute < 0 {
		hour--
		minute += 60
	}

	for hour < 0 {
		hour += 24
	}

	c.hour = (hour + minute/60) % 24
	c.minute = minute % 60
	return c
}

// String shows a time.
func (c Clock) String() string {
	switch {
	case c.hour/10 == 0 && c.minute/10 == 0:
		return "0" + strconv.Itoa(c.hour) + ":0" + strconv.Itoa(c.minute)
	case c.hour/10 == 0 && c.minute/10 != 0:
		return "0" + strconv.Itoa(c.hour) + ":" + strconv.Itoa(c.minute)
	case c.hour/10 != 0 && c.minute/10 == 0:
		return strconv.Itoa(c.hour) + ":0" + strconv.Itoa(c.minute)
	default:
		return strconv.Itoa(c.hour) + ":" + strconv.Itoa(c.minute)
	}
}

// Add adds the specified time.
func (c Clock) Add(minutes int) Clock {
	c.hour = (c.hour + (c.minute+minutes)/60) % 24
	c.minute = (c.minute + minutes) % 60
	return c
}

// Subtract subtracts the specified time.
func (c Clock) Subtract(minutes int) Clock {
	total := c.hour*60 + c.minute
	c.hour = (total - minutes) / 60
	c.minute = (total - minutes) % 60

	if c.minute < 0 {
		c.hour--
		c.minute += 60
	}

	for c.hour < 0 {
		c.hour += 24
	}

	return c
}
