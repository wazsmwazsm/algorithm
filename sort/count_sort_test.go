package sort

import (
	"testing"
)

func TestCountSort(t *testing.T) {
	a := []int{7, 5, 9, 11, 4, 4, 2, 6}
	res := []int{2, 4, 4, 5, 6, 7, 9, 11}
	CountSort(a)
	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Error("sort faild")
		}
	}
}

func TestCountSortStable(t *testing.T) {
	a := []int{7, 5, 9, 11, 4, 4, 2, 6}
	res := []int{2, 4, 4, 5, 6, 7, 9, 11}
	CountSortStable(a)
	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Error("sort faild")
		}
	}
}
