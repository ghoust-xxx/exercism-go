package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var cnt int
	for i := 0; i < len(birdsPerDay); i++ {
		cnt += birdsPerDay[i]
	}
	return cnt
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var cnt int
	for i := (week - 1) * 7; i < week * 7; i++ {
		cnt += birdsPerDay[i]
	}
	return cnt
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
