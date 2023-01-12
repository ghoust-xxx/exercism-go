package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings have different length")
	}
	ra, rb := []rune(a), []rune(b)
	cnt := 0
	for i := 0; i < len(ra); i++ {
		if ra[i] != rb[i] {
			cnt++
		}
	}
	return cnt, nil
}
