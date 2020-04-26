package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{7, 5, 9, 11, 4, 2, 6}
	res := []int{2, 4, 5, 6, 7, 9, 11}
	BubbleSort(a)
	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Error("sort faild")
		}
	}
}

func TestBubbleSort2(t *testing.T) {
	a := []int{1, 5, 9, 11, 4, 2, 6}
	res := []int{1, 2, 4, 5, 6, 9, 11}
	BubbleSort2(a)
	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Error("sort faild")
		}
	}
}

func TestBubbleSort3(t *testing.T) {
	a := []int{1, 5, 9, 11, 4, 2, 6, 17, 20, 33}
	res := []int{1, 2, 4, 5, 6, 9, 11, 17, 20, 33}
	BubbleSort3(a)
	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Error("sort faild")
		}
	}
}

func TestCockTailSort(t *testing.T) {
	a := []int{1, 5, 9, 11, 4, 2, 6, 17, 20, 33}
	res := []int{1, 2, 4, 5, 6, 9, 11, 17, 20, 33}
	CockTailSort(a)
	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Error("sort faild")
		}
	}
}
func TestCockTailSort2(t *testing.T) {
	a := []int{1, 5, 9, 11, 4, 2, 6, 17, 20, 33}
	res := []int{1, 2, 4, 5, 6, 9, 11, 17, 20, 33}
	CockTailSort2(a)
	for i := 0; i < len(res); i++ {
		if res[i] != a[i] {
			t.Error("sort faild")
		}
	}
}
