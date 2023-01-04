package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[*~=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	cnt := 0
	for _, line := range lines {
		if re.MatchString(line) {
			cnt++
		}
	}
	return cnt
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`\s+User\s+(\w+)\s+`)
	var out []string
	for _, line := range lines {
		res := re.FindStringSubmatch(line)
		if len(res) > 0 {
			out = append(out, "[USR] " + res[1] + " " + line)
		} else {
			out = append(out, line)
		}
	}
	return out
}
