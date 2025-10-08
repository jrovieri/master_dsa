package primefactors

import (
	isprime "github.com/jrovieri/master_dsa/is_prime"
)

// O(n logn)
func PrimeFactorsNaive(n int64) []int64 {

	var ans []int64

	var i int64 = 2
	for i < n {
		if isprime.IsPrime(i) {
			x := i
			for n%x == 0 {
				ans = append(ans, i)
				x *= i
			}
		}
		i++
	}
	return ans
}

func PrimeFactors(n int64) []int64 {

	var ans []int64

	if n <= 1 {
		return ans
	}

	for n%2 == 0 {
		ans = append(ans, 2)
		n /= 2
	}

	for n%3 == 0 {
		ans = append(ans, 3)
		n /= 3
	}

	var i int64 = 5
	for i*i <= n {

		for n%i == 0 {
			ans = append(ans, i)
			n /= i
		}

		for n%(i+2) == 0 {
			ans = append(ans, (i + 2))
			n /= (i + 2)
		}
		i += 6
	}

	if n > 3 {
		ans = append(ans, n)
	}
	return ans
}
