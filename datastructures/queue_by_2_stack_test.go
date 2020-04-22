package datastructures

import (
	"testing"
)

func TestTwoStackQueue(t *testing.T) {
	q := NewTwoStackQueue()

	if !q.Empty() {
		t.Error("queue not empty")
	}

	q.Offer(2)
	q.Offer(7)
	if q.Empty() {
		t.Error("queue is empty")
	}

	if q.Peek() != 2 {
		t.Error("queue peek get err")
	}
	if !q.Empty() {
		q.Pull()
		if q.Pull() != 7 {
			t.Error("queue pull get err")
		}
	}
}
