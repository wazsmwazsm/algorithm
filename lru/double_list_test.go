package lru

import (
	"testing"
)

func TestDoubleList(t *testing.T) {
	l := new(DoubleList)

	l.AddFirst(&Node{Key: 1, Value: 2})
	l.AddFirst(&Node{Key: 3, Value: 5})
	node := &Node{Key: 4, Value: 8}
	l.AddFirst(node)
	l.AddFirst(&Node{Key: 5, Value: 21})
	l.AddFirst(&Node{Key: 2, Value: 3})
	if l.Size() != 5 {
		t.Error("size err")
	}
	if v := l.RemoveLast(); v.Key != 1 || v.Value != 2 {
		t.Error("remove err")
	}
	l.Remove(node)
	if v := l.RemoveLast(); v.Key != 3 || v.Value != 5 {
		t.Error("remove err")
	}

	if l.Size() != 2 {
		t.Error("size err")
	}
	if v := l.RemoveFirst(); v.Key != 2 || v.Value != 3 {
		t.Error("remove err")
	}
	l.AddLast(&Node{Key: 88, Value: 88})
	if v := l.RemoveLast(); v.Key != 88 || v.Value != 88 {
		t.Error("remove err")
	}
	l.RemoveLast()
	if v := l.RemoveLast(); v != nil {
		t.Error("remove err")
	}
}
