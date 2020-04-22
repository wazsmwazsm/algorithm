package datastructures

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewMySliceStack()

	if !s.Empty() {
		t.Error("stack not empty")
	}

	s.Push(2)
	s.Push(7)
	if s.Empty() {
		t.Error("stack is empty")
	}

	if s.Top() != 7 {
		t.Error("stack top get err")
	}
	if !s.Empty() {
		s.Pop()
		if s.Pop() != 2 {
			t.Error("stack pop get err")
		}
	}

}
