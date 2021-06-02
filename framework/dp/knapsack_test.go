package dp

import (
	"testing"
)

func TestKnapsack(t *testing.T) {

	if 6 != knapsack([]int{2, 1, 3}, []int{4, 2, 3}, 4) {
		t.Error("knapsack err")
	}
	if 27 != knapsack([]int{5, 6, 3, 1}, []int{4, 6, 12, 9}, 12) {
		t.Error("knapsack err")
	}
}
