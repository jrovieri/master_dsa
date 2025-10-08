package isprime

import "testing"

func TestIsPrime1(t *testing.T) {

	var input int64 = 3331
	expected := true

	t.Run("naive=37", func(t *testing.T) {
		result := IsPrimeNaive(input)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})

	t.Run("efficient=37", func(t *testing.T) {
		result := IsPrime(input)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})
}

func TestIsPrime2(t *testing.T) {

	var input int64 = 3331
	expected := true

	t.Run("naive=3331", func(t *testing.T) {
		result := IsPrimeNaive(input)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})

	t.Run("efficient=3331", func(t *testing.T) {
		result := IsPrime(input)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})
}

func TestIsPrime3(t *testing.T) {

	var input int64 = 240
	expected := false

	t.Run("naive=240", func(t *testing.T) {
		result := IsPrimeNaive(input)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})

	t.Run("efficient=240", func(t *testing.T) {
		result := IsPrime(input)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})
}
