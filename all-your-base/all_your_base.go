package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	var num int

	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	for i, val := range inputDigits {
		if val < 0 || val >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		num += int(math.Pow(float64(inputBase), float64(len(inputDigits)-1-i))) * val
	}

	res := []int{}
	for num > 0 {
		res = append([]int{num % outputBase}, res...)
		num /= outputBase
	}
	if len(res) == 0 {
		return []int{0}, nil
	}

	return res, nil
}
