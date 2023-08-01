package perfect

import (
	"errors"
)

type Classification int

const (
	ClassificationDeficient = iota + 1
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("Only positive numbers")

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return 0, ErrOnlyPositive
	}

	var i, sum int64
	for i = 1; i <= n/2; i++ {
		if n % i == 0 {
			sum += i
		}
	}

	if sum == n {
		return ClassificationPerfect, nil
	} else if sum > n {
		return ClassificationAbundant, nil
	}

	return ClassificationDeficient, nil
}
