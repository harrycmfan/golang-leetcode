package data_structure

/*
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?

Example:

LRUCache cache = new LRUCache( 2 /* capacity );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
 */


type Node struct {
	next *Node
	pre *Node
	key int
	val int
}

type LRUCache struct {
	capacity int
	mapping map[int]*Node
	head *Node
	tail *Node
}


func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		mapping: map[int]*Node{},
		head: head,
		tail: tail,
	}
}


func (this *LRUCache) Get(key int) int {
	node, _ := this.mapping[key]
	if node == nil {
		return -1
	}
	if node.next != this.tail {
		node.pre.next = node.next
		node.next.pre = node.pre
		node.next = this.tail
		node.pre = this.tail.pre
		this.tail.pre.next = node
		this.tail.pre = node
	}
	return node.val
}


func (this *LRUCache) Put(key int, value int)  {
	node, _ := this.mapping[key]
	if node != nil {
		node.val = value
		_ = this.Get(key)
	} else {
		newNode := &Node{
			key: key,
			val: value,
			next: this.tail,
			pre: this.tail.pre,
		}
		this.tail.pre.next = newNode
		this.tail.pre = newNode
		this.mapping[key] = newNode
		if len(this.mapping) > this.capacity {
			//fmt.Println(this.mapping, this.head.next.key)
			delete(this.mapping, this.head.next.key)
			//fmt.Println(this.mapping)
			this.head.next.next.pre = this.head
			this.head.next = this.head.next.next
		}
	}
	//fmt.Println(this.mapping)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
