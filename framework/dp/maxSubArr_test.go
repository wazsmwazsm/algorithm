package dp

import (
	"testing"
)

func TestMaxSubArr(t *testing.T) {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if 6 != maxSubArr(a) {
		t.Error("max sub arr err")
	}
	a = []int{8, -13, 4, 5, -9, 6, 7}
	if 13 != maxSubArr(a) {
		t.Error("max sub arr err")
	}
}
