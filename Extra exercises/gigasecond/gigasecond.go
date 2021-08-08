// Package gigasecond implements a simple time conversion.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond implements a simple time conversion.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
