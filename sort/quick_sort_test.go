package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{7, 5, 9, 11, 4, 2, 6}
	res := []int{2, 4, 5, 6, 7, 9, 11}
	QuickSort(a, 0, len(a)-1)
	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Error("sort faild")
		}
	}
}

func TestQuickSort2(t *testing.T) {
	a := []int{7, 5, 9, 11, 4, 2, 6}
	res := []int{2, 4, 5, 6, 7, 9, 11}
	QuickSort2(a, 0, len(a)-1)
	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Error("sort faild")
		}
	}
}

func TestQuickSort3(t *testing.T) {
	a := []int{7, 5, 9, 11, 4, 2, 6}
	res := []int{2, 4, 5, 6, 7, 9, 11}
	QuickSort3(a, 0, len(a)-1)
	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Error("sort faild")
		}
	}
}
