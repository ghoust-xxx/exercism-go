package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	out, _ := time.Parse("1/2/2006 15:04:05", date)
	return out
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	old, _ := time.Parse("January 2, 2006 15:04:05", date)
	if old.Before(time.Now()) {
		return true
	} else {
		return false
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	app, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if app.Hour() >= 12 && app.Hour() < 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	app, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.",
		app.Weekday(),
		app.Month(),
		app.Day(),
		app.Year(),
		app.Hour(),
		app.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
