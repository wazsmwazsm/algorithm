package lru

// LRU Least Recently Used 一个 k\v 结构的最近最少使用缓存
type LRU struct {
	cache *DoubleList   /* 双向链表，O(1) 的链表移动 */
	m     map[int]*Node /* map, O(1) 的读取 */
	cap   int
}

// NewLRUCache 创建 LRU 缓存
func NewLRUCache(cap int) *LRU {
	return &LRU{
		cache: new(DoubleList),
		m:     make(map[int]*Node, cap),
		cap:   cap,
	}
}

// Get 获取缓存
func (l *LRU) Get(key int) int {

	if node, ok := l.m[key]; ok { // 存在
		// 数据提到链表头
		l.Put(key, node.Value)
		return node.Value
	}

	return -1
}

// Put 设置缓存
func (l *LRU) Put(key, value int) {

	if v, ok := l.m[key]; ok { // 已经存在的数据
		l.cache.Remove(v) // 删除旧节点
		// 放到头部
		v.Value = value // 更新值
		l.cache.AddFirst(v)
		return
	}

	// 不存在
	node := &Node{Key: key, Value: value}
	if l.cap == l.cache.Size() { // 满了
		last := l.cache.RemoveLast() // 删掉链表尾部
		delete(l.m, last.Key)        // 从 map 中删除
	}
	// 直接添加头部且存入 map
	l.cache.AddFirst(node)
	l.m[key] = node
}
