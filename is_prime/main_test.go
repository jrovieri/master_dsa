package main

import "testing"

func TestIsPrimeNaive(t *testing.T) {

	t.Run("Is 37 prime?", func(t *testing.T) {
		expected := true
		result := IsPrimeNaive(3331)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})

	t.Run("Is 1037 prime?", func(t *testing.T) {
		expected := true
		result := IsPrimeNaive(7919)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})

	t.Run("Is 240 prime?", func(t *testing.T) {
		expected := false
		result := IsPrimeNaive(240)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})
}

func TestIsPrime(t *testing.T) {

	t.Run("Is 37 prime?", func(t *testing.T) {
		expected := true
		result := IsPrime(3331)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})

	t.Run("Is 1037 prime?", func(t *testing.T) {
		expected := true
		result := IsPrime(7919)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})

	t.Run("Is 240 prime?", func(t *testing.T) {
		expected := false
		result := IsPrime(240)
		if result != expected {
			t.Errorf("expected: %t, result: %t", expected, result)
		}
	})
}
