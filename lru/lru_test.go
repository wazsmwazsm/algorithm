package lru

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	cache := NewLRUCache(5)

	cache.Put(1, 2)
	printCache(cache)
	cache.Put(4, 6)
	cache.Put(2, 8)
	cache.Put(7, 1)
	printCache(cache)

	if cache.Get(2) != 8 {
		t.Error("get cache error")
	}
	printCache(cache)
	if cache.Get(4) != 6 {
		t.Error("get cache error")
	}
	printCache(cache)
	if cache.Get(111) != -1 {
		t.Error("get cache error")
	}
	cache.Put(11, 3)
	printCache(cache)
	cache.Put(22, 5)
	printCache(cache)
	cache.Put(2, 15)
	printCache(cache)
	if cache.Get(4) != 6 {
		t.Error("get cache error")
	}
	printCache(cache)
}

// 打印 cache 链表的状态
func printCache(c *LRU) {
	fmt.Println()
	curr := c.cache.head
	for curr != nil {
		fmt.Printf(" %d:%d ", curr.Key, curr.Value)
		curr = curr.next
	}
}

func BenchmarkLRUGet(b *testing.B) {
	cache := NewLRUCache(100)
	for i := 0; i < 100; i++ {
		cache.Put(i, i)
	}

	for i := 0; i < b.N; i++ {
		cache.Get(i)
	}
}

func BenchmarkLRUPut(b *testing.B) {
	cache := NewLRUCache(100)

	for i := 0; i < b.N; i++ {
		cache.Put(i, i)
	}
}
