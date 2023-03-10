package sieve

func Sieve(limit int) (out []int) {
	if limit <= 0 {
		return nil
	}

	var tmp = make([]int, limit/2 + 1)
	tmp[0] = 2
	for i := 1; i < len(tmp); i++ {
		tmp[i] = i * 2 + 1
	}

	for i := 0; i < len(tmp); i++ {
		if tmp[i] == 0 {
			continue
		}
		for j := i + 1; j < len(tmp); j++ {
			if tmp[j] == 0 {
				continue
			}
			if tmp[j] % tmp[i] == 0 {
				tmp[j] = 0
			}
		}
	}

	for i := 0; i < len(tmp); i++ {
		if tmp[i] != 0 {
			out = append(out, tmp[i])
		}
	}
	if out[len(out)-1] > limit {
		out = out[:len(out)-1]
	}

	return out
}
