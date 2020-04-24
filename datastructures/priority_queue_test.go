package datastructures

import (
	// "fmt"
	"testing"
)

func TestMaxPriorityQueue(t *testing.T) {
	q := NewMaxPriorityQueue()

	q.Offer(1)
	q.Offer(4)
	q.Offer(3)
	q.Offer(8)
	q.Offer(5)
	if q.Peek() != 8 {
		t.Error("peek error")
	}

	res := []int{8, 5, 4, 3, 1}
	for i := 0; i < len(res); i++ {
		if q.Len() != len(res)-i {
			t.Error("len error")
		}
		if q.Pull() != res[i] {
			t.Error("pull error")
		}
	}

}

func TestMinPriorityQueue(t *testing.T) {
	q := NewMinPriorityQueue()

	q.Offer(1)
	q.Offer(4)
	q.Offer(3)
	q.Offer(8)
	q.Offer(5)
	if q.Peek() != 1 {
		t.Error("peek error")
	}

	res := []int{1, 3, 4, 5, 8}
	for i := 0; i < len(res); i++ {
		if q.Len() != len(res)-i {
			t.Error("len error")
		}
		if q.Pull() != res[i] {
			t.Error("pull error")
		}

	}

}
