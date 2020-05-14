package lru

// 非并发安全, 实现各个接口 O(1) 操作

// Node 链表节点
type Node struct {
	Key, Value int /* 为了最终 LRU map 可以删除末尾元素而保存 key */
	prev       *Node
	next       *Node
}

// DoubleList 双向链表
type DoubleList struct {
	length int
	head   *Node
	tail   *Node
}

// AddFirst 添加元素到链表头
func (l *DoubleList) AddFirst(x *Node) {
	defer func() {
		l.length++
	}()

	if l.head == nil {
		l.head, l.tail = x, x
		return
	}
	l.head.prev = x            // 旧头 prev 链到新头
	l.head, x.next = x, l.head // 变为新头
}

// AddLast 添加元素到链表尾部
func (l *DoubleList) AddLast(x *Node) {
	defer func() {
		l.length++
	}()

	if l.head == nil {
		l.head, l.tail = x, x
		return
	}
	l.tail.next = x            // 旧尾 next 链到新头
	l.tail, x.prev = x, l.tail // 变为新尾
}

// Remove 删掉一个节点
func (l *DoubleList) Remove(x *Node) {
	defer func() {
		l.length--
	}()

	if l.length <= 0 {
		return
	}

	prev := x.prev
	next := x.next

	if prev == nil && next == nil { // 不在链上
		return
	}

	if prev == nil { // 链表头
		l.head, next.prev = next, nil
		return
	}
	if next == nil { // 链表尾
		l.tail, prev.next = prev, nil
		return
	}

	// 中间
	prev.next, next.prev = next, prev
}

// RemoveLast 删掉末尾节点
func (l *DoubleList) RemoveLast() *Node {
	defer func() {
		l.length--
	}()

	if l.length <= 0 {
		return nil
	}

	tail := l.tail
	prev := l.tail.prev
	if prev == nil { // 就一个元素
		l.head, l.tail = nil, nil
		return tail
	}

	l.tail, prev.next = prev, nil
	return tail
}

// RemoveFirst 删掉开头节点
func (l *DoubleList) RemoveFirst() *Node {
	defer func() {
		l.length--
	}()

	if l.length <= 0 {
		return nil
	}

	head := l.head
	next := l.head.next
	if next == nil { // 就一个元素
		l.head, l.tail = nil, nil
		return head
	}

	l.head, next.prev = next, nil
	return head
}

// Size 获取链表长度
func (l *DoubleList) Size() int {
	return l.length
}
