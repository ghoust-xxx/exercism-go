package prime

import (
	"errors"
	"math"
)

var primes []bool

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("there is no zeroth prime")
	}

	if primes == nil {
		getPrimes()
	}

	cnt := 1
	for i := 2; i < len(primes); i++ {
		if !primes[i] {
			continue
		}
		if cnt == n {
			return i, nil
		}
		cnt++
	}

	return 1, nil
}

func getPrimes() {
	var j int
	primes = make([]bool, math.MaxInt32/1000)
	primes[2] = true
	for i := 3; i < len(primes); i += 2 {
		primes[i] = true
	}

	for i := 3; i < len(primes); i++ {
		if !primes[i] {
			continue
		}
		j = 2 * i
		for j < len(primes) {
			primes[j] = false
			j += i
		}
	}
}

