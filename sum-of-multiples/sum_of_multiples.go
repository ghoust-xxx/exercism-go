package summultiples

func SumMultiples(limit int, divisors ...int) int {
	res := 0
	nums := make(map[int]bool)
	for _, div := range divisors {
		if div == 0 {
			continue
		}
		for i := 1; div * i < limit; i++ {
			nums[div * i] = true
		}
	}
	for key, _ := range nums {
		res += key
	}

	return res
}
