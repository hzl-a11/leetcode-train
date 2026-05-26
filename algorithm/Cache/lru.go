package algorithm

// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
// https://leetcode.cn/problems/lru-cache/submissions/727285614/?envType=problem-list-v2&envId=xitongyuruanjiansheji-sd1-huancunxitongsheji
type Node struct {
	// 双向链表，保证移动O1
	k, v      int
	pre, next *Node
}
type LRUCache struct {
	capacity   int           //优先容量
	m          map[int]*Node //缓存map，保证查询O1
	head, tail *Node         //链表的头尾指针,应为虚拟节点，以方便增删
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		capacity: capacity,
		head:     &Node{},
		tail:     &Node{},
		m:        make(map[int]*Node),
	}
	c.head.next = c.tail
	c.tail.pre = c.head
	return c
}

func (this *LRUCache) deleteNode(node *Node) {
	if node == nil {
		return
	}
	nextNode := node.next
	if node.pre != nil {
		node.pre.next = nextNode
	}

	nextNode.pre = node.pre

}
func (this *LRUCache) insertToHead(node *Node) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node

}
func (this *LRUCache) moveNodeToHead(node *Node) {
	this.deleteNode(node)
	this.insertToHead(node)
}

func (this *LRUCache) removeLastNode() {
	lastNode := this.tail.pre
	this.deleteNode(lastNode)
	delete(this.m, lastNode.k)
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if !ok {
		return -1
	}
	//把节点移动到链表头
	this.moveNodeToHead(node)
	return node.v
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.m[key]
	if ok {
		node.v = value
		this.moveNodeToHead(node)
	} else {
		if len(this.m) == this.capacity {
			this.removeLastNode()
		}
		node = &Node{k: key, v: value}
		this.m[key] = node
		this.insertToHead(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
