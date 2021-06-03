package dp

import (
	"testing"
)

func TestCanPartition(t *testing.T) {

	if !canPartition([]int{1, 5, 11, 5}) {
		t.Error("canPartition err")
	}
	if canPartition([]int{1, 2, 3, 5}) {
		t.Error("canPartition err")
	}
	if canPartition([]int{7, 9, 8, 4}) {
		t.Error("canPartition err")
	}
}
func TestCanPartition2(t *testing.T) {

	if !canPartition2([]int{1, 5, 11, 5}) {
		t.Error("canPartition err")
	}
	if canPartition2([]int{1, 2, 3, 5}) {
		t.Error("canPartition err")
	}
	if canPartition2([]int{7, 9, 8, 4}) {
		t.Error("canPartition err")
	}
}
