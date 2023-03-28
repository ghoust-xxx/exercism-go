package series

func All(n int, s string) []string {
	var res []string

	for i := 0; i + n <= len(s); i++ {
		res = append(res, s[i:i+n])
	}

	return res
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if len(s) < n {
		return "", false
	}

	return UnsafeFirst(n, s), true
}
