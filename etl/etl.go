package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
	for i, lts := range in {
		for _, ch := range lts {
			out[strings.ToLower(ch)] = i
		}
	}
	return out
}
