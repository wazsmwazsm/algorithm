package datastructures

import (
	"fmt"
	"testing"
)

func TestBuildMaxHeap(t *testing.T) {
	mh := NewMaxHeap(1, 5, 4, 3, 6, 7, 2)

	fmt.Println("init max head:", mh)
}
func TestMaxHeap(t *testing.T) {
	mh := NewMaxHeap()

	fmt.Println("init max head:", mh)
	mh.Insert(5)
	fmt.Println("insert 5 now max head:", mh)
	mh.Insert(6)
	fmt.Println("insert 6 now max head:", mh)
	mh.Insert(4)
	fmt.Println("insert 4 now max head:", mh)
	mh.Insert(8)
	fmt.Println("insert 8 now max head:", mh)
	mh.Insert(7)
	fmt.Println("insert 2 now max head:", mh)

	fmt.Println("Top is:", mh.Top())

	fmt.Printf("delete heap top %d now max head:%v\n", mh.DeleteMax(), mh)
	fmt.Printf("delete heap top %d now max head:%v\n", mh.DeleteMax(), mh)
	fmt.Printf("delete heap top %d now max head:%v\n", mh.DeleteMax(), mh)
	fmt.Printf("delete heap top %d now max head:%v\n", mh.DeleteMax(), mh)
	fmt.Printf("delete heap top %d now max head:%v\n", mh.DeleteMax(), mh)
	fmt.Printf("delete heap top %d now max head:%v\n", mh.DeleteMax(), mh)
}
