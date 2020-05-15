package lru

import (
	"fmt"
	"testing"
)

func TestLRU2(t *testing.T) {
	cache := NewLRU2Cache(5)

	cache.Put(1, 2)
	printCache2(cache)
	cache.Put(4, 6)
	cache.Put(2, 8)
	cache.Put(7, 1)
	printCache2(cache)

	if cache.Get(2) != 8 {
		t.Error("get cache error")
	}
	printCache2(cache)
	if cache.Get(4) != 6 {
		t.Error("get cache error")
	}
	printCache2(cache)
	if cache.Get(111) != -1 {
		t.Error("get cache error")
	}
	cache.Put(11, 3)
	printCache2(cache)
	cache.Put(22, 5)
	printCache2(cache)
	cache.Put(2, 15)
	printCache2(cache)
	if cache.Get(4) != 6 {
		t.Error("get cache error")
	}
	printCache2(cache)
}

// 打印 cache 链表的状态
func printCache2(c *LRU2) {
	fmt.Println()
	curr := c.cache.Front()
	for curr != nil {
		fmt.Printf(" %d:%d ", curr.Value.(kv).key, curr.Value.(kv).value)
		curr = curr.Next()
	}
}

func BenchmarkLRU2Get(b *testing.B) {
	cache := NewLRUCache(100)
	for i := 0; i < 100; i++ {
		cache.Put(i, i)
	}

	for i := 0; i < b.N; i++ {
		cache.Get(i)
	}
}

func BenchmarkLRU2Put(b *testing.B) {
	cache := NewLRUCache(100)

	for i := 0; i < b.N; i++ {
		cache.Put(i, i)
	}
}
