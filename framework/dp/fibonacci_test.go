package dp

import (
	"testing"
)

func TestFibonacciRecursion(t *testing.T) {
	res := []int{1, 1, 2, 3, 5, 8, 13, 21}

	for i := 0; i < len(res); i++ {
		if res[i] != fibonacciRecursion(i+1) {
			t.Error("fibo err")
		}
	}
}

func TestFibonacciRecursionWithMem(t *testing.T) {
	res := []int{1, 1, 2, 3, 5, 8, 13, 21}

	for i := 0; i < len(res); i++ {
		if res[i] != fibonacciRecursionWithMem(i+1) {
			t.Error("fibo err")
		}
	}
}

func TestFibonacci(t *testing.T) {
	res := []int{1, 1, 2, 3, 5, 8, 13, 21}

	for i := 0; i < len(res); i++ {
		if res[i] != fibonacci(i+1) {
			t.Error("fibo err")
		}
	}
}

func TestFibonacci2(t *testing.T) {
	res := []int{1, 1, 2, 3, 5, 8, 13, 21}

	for i := 0; i < len(res); i++ {
		if res[i] != fibonacci2(i+1) {
			t.Error("fibo err")
		}
	}
}
