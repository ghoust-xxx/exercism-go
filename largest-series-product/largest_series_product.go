package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var max int64
	if span > len(digits) {
		return 0, errors.New("not enought digits")
	}
	if span < 0 {
		return 0, errors.New("rejects negative span")
	}
	max = 0
	for i := 0; i < len(digits) - span + 1; i++ {
		curr := series(digits[i:i + span])
		if curr < 0 {
			return 0, errors.New("digits input must only contain digits")
		}
		if curr > max {
			max = curr
		}
	}

	return max, nil
}

func series(str string) (res int64) {
	res = 1
	for _, val := range str {
		num, err := strconv.Atoi(string(val))
		if err != nil {
			return -1
		}
		res *= int64(num)
	}

	return res
}
