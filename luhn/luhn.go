package luhn

import "strconv"

func Valid(id string) bool {
	slen := 0
	even := true
	crc := 0
	for i := len(id) - 1; i >= 0; i-- {
		if string(id[i]) == " " {
			continue
		}
		var num int
		var err error
		if num, err = strconv.Atoi(string(id[i])); err != nil {
			return false
		}
		slen++
		even = !even
		if !even {
			crc += num
			continue
		}
		num += num
		if num > 9 {
			num -= 9
		}
		crc += num
	}

	if slen <= 1 {
		return false
	}

	return crc % 10 == 0
}
