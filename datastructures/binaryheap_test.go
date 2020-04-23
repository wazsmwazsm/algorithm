package datastructures

import (
	"fmt"
	"testing"
)

func TestMinHead(t *testing.T) {
	mh := NewMinHead()

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

	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
	fmt.Printf("delete heap top %d now min head:%v\n", mh.DeleteMin(), mh)
}
