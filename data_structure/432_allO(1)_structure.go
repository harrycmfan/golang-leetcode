package data_structure

/*
Implement a data structure supporting the following operations:

Inc(Key) - Inserts a new key with value 1. Or increments an existing key by 1. Key is guaranteed to be a non-empty string.
Dec(Key) - If Key's value is 1, remove it from the data structure. Otherwise decrements an existing key by 1. If the key does not exist, this function does nothing. Key is guaranteed to be a non-empty string.
GetMaxKey() - Returns one of the keys with maximal value. If no element exists, return an empty string "".
GetMinKey() - Returns one of the keys with minimal value. If no element exists, return an empty string "".
Challenge: Perform all these in O(1) time complexity.
 */

type Node struct {
	val int
	keys map[string]bool
	next *Node
	pre *Node
}

type AllOne struct {
	mapping map[string]*Node
	head *Node
	tail *Node
}


/** Initialize your data structure here. */
func Constructor() AllOne {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return AllOne{
		mapping: map[string]*Node{},
		head: head,
		tail: tail,
	}
}


/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string)  {
	node, _ := this.mapping[key]
	if node == nil { // key doesn't exist
		if this.head.next == this.tail { // first element
			newNode := &Node{
				val: 1,
				keys: map[string]bool{key:true},
			}
			this.head.next = newNode
			newNode.pre = this.head
			this.tail.pre = newNode
			newNode.next = this.tail
			this.mapping[key] = newNode
		} else if this.head.next.val != 1 { // 1 doesn't exist
			newNode := &Node{
				val: 1,
				keys: map[string]bool{key:true},
			}
			jumpNode := this.head.next
			this.head.next = newNode
			newNode.pre = this.head
			jumpNode.pre = newNode
			newNode.next = jumpNode
			this.mapping[key] = newNode
		} else { // 1 exits
			this.head.next.keys[key] = true
			this.mapping[key] = this.head.next
		}
	} else { // key exists there
		if len(node.keys) == 1 { // node has itself only
			if node.next == this.tail || node.next.val != node.val + 1 { // node is max or node next is not adjacent node
				node.val++
			} else { // node has a next node which need to merge
				mergeNode := node.next
				mergeNode.keys[key] = true
				mergeNode.pre = node.pre
				node.pre.next = mergeNode
				this.mapping[key] = mergeNode
			}
		} else { // node has other key
			if node.next == this.tail || node.next.val != node.val + 1 { // node is max or node next is not adjacent node
				newNode := &Node{
					val: node.val + 1,
					keys: map[string]bool{key:true},
				}
				nextNode := node.next
				node.next = newNode
				newNode.pre = node
				nextNode.pre = newNode
				newNode.next = nextNode
				delete(node.keys, key)
				this.mapping[key] = newNode
			} else { // node has a next node so just move the key
				mergeNode := node.next
				mergeNode.keys[key] = true
				delete(node.keys, key)
				this.mapping[key] = mergeNode
			}
		}
	}
	// next := this.head.next
	// for next != this.tail {
	//     fmt.Print(next.val, next.keys)
	//     next = next.next
	// }
	// fmt.Println()
}


/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {
	node, _ := this.mapping[key]
	if node == nil { // key doesn't exist
		return
	}
	if node.val == 1 { // node is 1
		if len(node.keys) == 1 { // only him
			nextNode := node.next
			this.head.next = nextNode
			nextNode.pre = this.head
			delete(this.mapping, key)
		} else { // has others, then just remove him
			delete(node.keys, key)
			delete(this.mapping, key)
		}
	} else { // node is not 1
		if len(node.keys) == 1 { // only him
			if node.pre == this.head || node.pre.val != node.val - 1 { // node is head or node pre is not a adjacent node
				node.val--
			} else { // node need to merge to small one
				mergeNode := node.pre
				mergeNode.keys[key] = true
				mergeNode.next = node.next
				node.next.pre = mergeNode
				this.mapping[key] = mergeNode
			}
		} else { // not only him
			if node.pre == this.head || node.pre.val != node.val - 1 { // node is head or node pre is not a adjacent node
				newNode := &Node{
					val: node.val - 1,
					keys: map[string]bool{key:true},
				}
				preNode := node.pre
				node.pre = newNode
				newNode.next = node
				preNode.next = newNode
				newNode.pre = preNode
				delete(node.keys, key)
				this.mapping[key] = newNode
			} else { // just need to merge
				mergeNode := node.pre
				mergeNode.keys[key] = true
				delete(node.keys, key)
				this.mapping[key] = mergeNode
			}
		}
	}
}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.tail.pre != this.head {
		for k := range this.tail.pre.keys {
			return k
		}
	}
	return ""
}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.head.next != this.tail {
		for k := range this.head.next.keys {
			return k
		}
	}
	return ""
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */