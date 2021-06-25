package dp

import (
	"testing"
)

func TestCoinChangeKnapsack(t *testing.T) {
	res := []int{0, 1, 2, 2, 3, 4}
	coins := []int{1, 2, 5}

	for i := 0; i < len(res); i++ {

		if res[i] != change(i, coins) {
			t.Error("coin change err")
		}
	}
}
func TestCoinChangeKnapsack2(t *testing.T) {
	res := []int{0, 1, 2, 2, 3, 4}
	coins := []int{1, 2, 5}

	for i := 0; i < len(res); i++ {

		if res[i] != change2(i, coins) {
			t.Error("coin change err")
		}
	}
}
func TestCoinChangeKnapsack3(t *testing.T) {
	res := []int{0, 1, 2, 2, 3, 4}
	coins := []int{1, 2, 5}

	for i := 0; i < len(res); i++ {

		if res[i] != change3(i, coins) {
			t.Error("coin change err")
		}
	}
}
