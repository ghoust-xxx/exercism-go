package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	if strings.HasPrefix(question, "What is ") {
		question = question[8:]
	} else {
		return 0, false
	}
	if strings.HasSuffix(question, "?") {
		question = question[:len(question)-1]
	}

	words := strings.Split(question, " ")
	res, err := strconv.Atoi(words[0])
	if err != nil {
		return 0, false
	}
	words = words[1:]
	for len(words) > 0 {
		switch words[0] {
		case "plus":
			if len(words) < 2 {
				return 0, false
			}
			op, err := strconv.Atoi(words[1])
			if err != nil {
				return 0, false
			}
			res += op
			words = words[2:]
		case "minus":
			if len(words) < 2 {
				return 0, false
			}
			op, err := strconv.Atoi(words[1])
			if err != nil {
				return 0, false
			}
			res -= op
			words = words[2:]
		case "multiplied":
			if len(words) < 3 || words[1] != "by" {
				return 0, false
			}
			op, err := strconv.Atoi(words[2])
			if err != nil {
				return 0, false
			}
			res *= op
			words = words[3:]
		case "divided":
			if len(words) < 3 || words[1] != "by" {
				return 0, false
			}
			op, err := strconv.Atoi(words[2])
			if err != nil {
				return 0, false
			}
			res /= op
			words = words[3:]
		default:
			return 0, false
		}
	}

	return res, true
}
