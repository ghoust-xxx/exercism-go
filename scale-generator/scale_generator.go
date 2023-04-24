package scale

import (
	"unicode"
)

var notes_sharp []string = []string{
	"A",
	"A#",
	"B",
	"C",
	"C#",
	"D",
	"D#",
	"E",
	"F",
	"F#",
	"G",
	"G#",
}
var notes_flat []string = []string{
	"A",
	"Bb",
	"B",
	"C",
	"Db",
	"D",
	"Eb",
	"E",
	"F",
	"Gb",
	"G",
	"Ab",
}

func Scale(tonic, interval string) []string {
	var res []string
	var notes []string

	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		notes = notes_sharp
	default:
		notes = notes_flat
	}

	r := []rune(tonic)
	r[0] = unicode.ToUpper(r[0])
	tonic = string(r)

	var idx int
	for i, val := range notes {
		if val == tonic {
			idx = i
			break
		}
	}

	if interval == "" {
		for i := 0; i < 12; i++ {
			res = append(res, notes[(i + idx) % 12])
		}
		return res
	}

	cnt := 0
	for _, val := range interval {
		res = append(res, notes[(idx + cnt) % 12])
		if val == 'm' {
			cnt++
		} else if val == 'M' {
			cnt += 2
		} else {
			cnt += 3
		}
	}
	res = append(res, notes[(idx + cnt) % 12])

	return res
}
