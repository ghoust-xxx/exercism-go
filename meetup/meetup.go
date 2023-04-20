package meetup

import "time"

type WeekSchedule int
const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Fifth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	wd := t.Weekday()
	diff := (int(wDay - wd) + 7) % 7 + 1 + 7 * int(wSched)
	switch wSched {
	case First:
		return diff
	case Second:
		return diff
	case Third:
		return diff
	case Fourth:
		return diff
	case Fifth:
		return diff
	case Last:
		t = t.AddDate(0, 1, -1)
		wd := t.Weekday()
		return t.Day() - (int(wd - wDay) + 7) % 7
	case Teenth:
		t = time.Date(year, month, 13, 0, 0, 0, 0, time.Local)
		wd := t.Weekday()
		return t.Day() + (int(wDay - wd) + 7) % 7
	default:
		return 0
	}

	return 42
}
