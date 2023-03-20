package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Duration(1*1000*1000*1000*time.Second))
	return t
}
