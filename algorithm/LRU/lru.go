// 146. LRU 缓存
// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
// https://leetcode.cn/problems/lru-cache/

package lru

// 双向链表方便增删，map快速判断是否存在以及O（1）的get。注意双向链表的头尾指针最好用虚拟指针，可以有效避免增删时大量的异常处理

type Node struct {
	key, val  int
	pre, next *Node
}

type LRUCache struct {
	capacity int
	m        map[int]*Node
	// 哨兵：head.next 指向最久未使用的(LRU)，tail.pre 指向最近使用的(MRU)
	// 提示：你可以反过来定义，只要逻辑统一即可
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		capacity: capacity,
		m:        make(map[int]*Node),
		head:     &Node{}, // 虚拟头
		tail:     &Node{}, // 虚拟尾
	}
	c.head.next = c.tail
	c.tail.pre = c.head
	return c
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.moveToTail(node) // 访问后，移动到尾部表示“最近使用”
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.val = value
		this.moveToTail(node)
		return
	}

	if len(this.m) >= this.capacity {
		// 删掉最久没用的：head.next 指向的那个
		oldest := this.head.next
		this.removeNode(oldest)
		delete(this.m, oldest.key)
	}

	newNode := &Node{key: key, val: value}
	this.m[key] = newNode
	this.addToTail(newNode)
}

// --- 以下是极度简化的链表操作，再也不用担心 nil ---

func (this *LRUCache) removeNode(n *Node) {
	n.pre.next = n.next
	n.next.pre = n.pre
}

func (this *LRUCache) addToTail(n *Node) {
	// 插入到 tail 和 tail.pre 之间
	n.next = this.tail
	n.pre = this.tail.pre
	this.tail.pre.next = n
	this.tail.pre = n
}

func (this *LRUCache) moveToTail(n *Node) {
	this.removeNode(n)
	this.addToTail(n)
}
