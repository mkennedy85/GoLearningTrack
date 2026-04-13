// Package gigasecond determine the date and time one gigasecond after a certain date.
package gigasecond

import "time"

// AddGigasecond returns a Time representing the input Time with a gigasecond added to it.
func AddGigasecond(t time.Time) time.Time {
	gigaTime := t.Add(time.Second * 1000000000)
	return gigaTime
}
