package datastructures

import (
	"testing"
)

func TestTwoStackQueue(t *testing.T) {
	q := NewTwoStackQueue()

	if !q.Empty() {
		t.Error("queue not empty")
	}

	q.Push(2)
	q.Push(7)
	if q.Empty() {
		t.Error("queue is empty")
	}

	if q.Peek() != 2 {
		t.Error("queue peek get err")
	}
	if !q.Empty() {
		q.Pop()
		if q.Pop() != 7 {
			t.Error("queue pop get err")
		}
	}
}
