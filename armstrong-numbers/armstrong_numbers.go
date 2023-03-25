package armstrong

import (
	"strconv"
)

func IsNumber(n int) bool {
	str := strconv.Itoa(n)
	sum := 0
	for _, val := range str {
		n, _ := strconv.Atoi(string(val))
		mul := 1
		for i := 0; i < len(str); i++ {
			mul *= n
		}
		sum += mul
	}

	return sum == n
}
