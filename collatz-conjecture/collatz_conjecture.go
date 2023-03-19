package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	cnt := 0
	if n < 1 {
		return 0, errors.New("number should be positive")
	}

	for n != 1 {
		cnt++
		if n % 2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
	}

	return cnt, nil
}
