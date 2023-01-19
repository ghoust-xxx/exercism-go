package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("error")
	}
	return 1 << (number - 1), nil
}

func Total() uint64 {
	last, _ := Square(64)
	return (last << 1) - 1
}
