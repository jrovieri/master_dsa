package main

func IsPrimeNaive(n int64) bool {

	if n == 1 {
		return false
	}

	var c int64 = 2
	for c < n {
		if n%c == 0 {
			return false
		}
		c++
	}
	return true
}

func IsPrime(n int64) bool {

	if n == 1 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	var c int64 = 5
	for c*c <= n {
		if n%c == 0 || n%(c+2) == 0 {
			return false
		}
		c += 6
	}
	return true
}
