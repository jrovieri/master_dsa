package primefactors

import (
	"slices"
	"testing"
)

func TestPrimeFactors1(t *testing.T) {

	var input int64 = 6
	expected := []int64{2, 3}

	t.Run("naive=6", func(t *testing.T) {
		result := PrimeFactorsNaive(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})

	t.Run("efficient=6", func(t *testing.T) {
		result := PrimeFactors(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})
}

func TestPrimeFactors2(t *testing.T) {

	var input int64 = 12
	expected := []int64{2, 2, 3}

	t.Run("naive=12", func(t *testing.T) {
		result := PrimeFactorsNaive(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})

	t.Run("efficient=12", func(t *testing.T) {
		result := PrimeFactors(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})
}

func TestPrimeFactors3(t *testing.T) {

	var input int64 = 84
	expected := []int64{2, 2, 3, 7}

	t.Run("naive=84", func(t *testing.T) {
		result := PrimeFactorsNaive(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})

	t.Run("efficient=84", func(t *testing.T) {
		result := PrimeFactors(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})
}

func TestPrimeFactors4(t *testing.T) {

	var input int64 = 132
	expected := []int64{2, 2, 3, 11}

	t.Run("naive=132", func(t *testing.T) {
		result := PrimeFactorsNaive(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})

	t.Run("efficient=132", func(t *testing.T) {
		result := PrimeFactors(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})
}

func TestPrimeFactors5(t *testing.T) {

	var input int64 = 450
	expected := []int64{2, 3, 3, 5, 5}

	t.Run("naive=450", func(t *testing.T) {
		result := PrimeFactorsNaive(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})

	t.Run("efficient=450", func(t *testing.T) {
		result := PrimeFactors(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})
}

func TestPrimeFactors6(t *testing.T) {

	var input int64 = 1027
	expected := []int64{13, 79}

	t.Run("naive=1027", func(t *testing.T) {
		result := PrimeFactorsNaive(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})

	t.Run("efficient=1027", func(t *testing.T) {
		result := PrimeFactors(input)
		if !slices.Equal(expected, result) {
			t.Errorf("expected: %v, result: %v", expected, result)
		}
	})
}
