package main

import (
	"errors"
	"math"
)

// Part-1
// Factorial calculates n!
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	if n == 0 {
		return 1, nil
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// IsPrime returns true if n is a prime number
func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}
	if n == 2 {
		return true, nil
	}
	if n%2 == 0 {
		return false, nil
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

// Power calculates base^exponent
func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result, nil
}

// Part-2
// MakeCounter
func MakeCounter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

// MakeMultiplier
func MakeMultiplier(factor int) func(int) int {
	return func(input int) int {
		return input * factor
	}
}

// MakeAccumulator
func MakeAccumulator(initial int) (add func(int), subtract func(int), get func() int) {
	state := initial

	add = func(n int) {
		state += n
	}
	subtract = func(n int) {
		state -= n
	}
	get = func() int {
		return state
	}

	return add, subtract, get
}

func main() {

}
