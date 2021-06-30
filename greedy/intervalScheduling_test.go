package greedy

import (
	"testing"
)

func TestIntervalScheduling(t *testing.T) {

	if 2 != intervalScheduling([][]int{{1, 3}, {2, 4}, {3, 6}}) {
		t.Error("interval scheduling err")
	}

	if 3 != intervalScheduling([][]int{{1, 2}, {2, 6}, {7, 9}, {3, 8}}) {
		t.Error("interval scheduling err")
	}
}
