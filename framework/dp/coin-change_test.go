package dp

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	res := []int{0, 1, 1, 2, 2, 1, 2, 2, 3, 3, 2}
	coins := []int{1, 2, 5}

	for i := 0; i < len(res); i++ {

		if res[i] != coinChange(coins, i) {
			t.Error("coin change err")
		}
	}
}

func TestCoinChange3(t *testing.T) {
	res := []int{0, 1, 1, 2, 2, 1, 2, 2, 3, 3, 2}
	coins := []int{1, 2, 5}

	for i := 0; i < len(res); i++ {

		if res[i] != coinChange3(coins, i) {
			t.Error("coin change err")
		}
	}
}
