package lru

import (
	"container/list"
)

// 存放 kv 数据
type kv struct {
	key, value int
}

// LRU2  使用 go 提供的双向链表来做
type LRU2 struct {
	cache *list.List            /* 双向链表，O(1) 的链表移动 */
	m     map[int]*list.Element /* map, O(1) 的读取 */
	cap   int
}

// NewLRU2Cache 创建 LRU 缓存
func NewLRU2Cache(cap int) *LRU2 {
	return &LRU2{
		cache: list.New(),
		m:     make(map[int]*list.Element, cap),
		cap:   cap,
	}
}

// Get 获取缓存
func (l *LRU2) Get(key int) int {

	if el, ok := l.m[key]; ok { // 存在
		value := el.Value.(kv).value
		// 数据提到链表头
		l.Put(key, value)
		return value
	}

	return -1
}

// Put 设置缓存
func (l *LRU2) Put(key, value int) {

	if v, ok := l.m[key]; ok { // 已经存在的数据
		l.cache.Remove(v) // 删除旧节点
		// 新值放到头部
		l.cache.PushFront(kv{key: key, value: value})
		return
	}

	// 不存在
	v := kv{key: key, value: value}
	if l.cap == l.cache.Len() { // 满了
		last := l.cache.Remove(l.cache.Back()) // 删掉链表尾部
		delete(l.m, last.(kv).key)             // 从 map 中删除
	}
	// 直接添加头部且存入 map
	l.m[key] = l.cache.PushFront(v)
}
