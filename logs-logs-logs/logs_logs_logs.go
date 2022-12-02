package logs

import (
	"bytes"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, val := range log {
		switch val {
			case '\U00002757':
				return "recommendation"
			case '\U0001F50D':
				return "search"
			case '\U00002600':
				return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var buf bytes.Buffer
	for _, val := range log {
		if val == oldRune {
			buf.WriteString(string(newRune))
		} else {
			buf.WriteString(string(val))
		}
	}
	return buf.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
