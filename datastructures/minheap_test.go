package datastructures

import (
	"fmt"
	"testing"
)

func TestBuildMinHeap(t *testing.T) {
	mh := NewMinHeap(1, 5, 4, 3, 6, 7, 2)

	fmt.Println("init max head:", mh)
}
func TestMinHeap(t *testing.T) {
	mh := NewMinHeap()

	fmt.Println("init min head:", mh)
	mh.Insert(5)
	fmt.Println("insert 5 now min head:", mh)
	mh.Insert(6)
	fmt.Println("insert 6 now min head:", mh)
	mh.Insert(4)
	fmt.Println("insert 4 now min head:", mh)
	mh.Insert(3)
	fmt.Println("insert 3 now min head:", mh)
	mh.Insert(2)
	fmt.Println("insert 2 now min head:", mh)

	fmt.Println("Top is:", mh.Top())

	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
}
