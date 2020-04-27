package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	a := []int{6, 4, 2, 7, 8, 9, 1}
	res := []int{1, 2, 4, 6, 7, 8, 9}
	HeapSort(a)
	for i := 0; i < len(a); i++ {
		if a[i] != res[i] {
			t.Error("sort faild")
		}
	}
	fmt.Println(a) // [1 2 4 6 7 8 9]
}

func TestHeapSort2(t *testing.T) {
	a := []int{6, 4, 2, 7, 8, 9, 1}
	res := []int{1, 2, 4, 6, 7, 8, 9}
	HeapSort2(a)
	for i := 0; i < len(a); i++ {
		if a[i] != res[i] {
			t.Error("sort faild")
		}
	}
	fmt.Println(a) // [1 2 4 6 7 8 9]
}

func TestHeapSort3(t *testing.T) {
	a := []int{6, 4, 2, 7, 8, 9, 1}
	res := []int{1, 2, 4, 6, 7, 8, 9}
	HeapSort3(a)
	for i := 0; i < len(a); i++ {
		if a[i] != res[i] {
			t.Error("sort faild")
		}
	}
	fmt.Println(a) // [1 2 4 6 7 8 9]
}
