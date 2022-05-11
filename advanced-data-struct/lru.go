// LC146. LRU 缓存
package main

import "fmt"

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func printList(head *Node) {
	if d {
		fmt.Printf("list: ")
		for head != nil {
			fmt.Printf("[%v %v] --> ", head.Key, head.Val)
			head = head.Back
		}
		fmt.Printf("nil\n")
	}
}

type Node struct {
	Key   int
	Val   int
	Front *Node
	Back  *Node
}

type LRUCache struct {
	Map   map[int]*Node
	Head  *Node
	Tail  *Node
	Count int
	Cap   int
}

func Constructor(capacity int) LRUCache {
	var lru = LRUCache{Head: &Node{}, Tail: &Node{}}
	lru.Head.Back = lru.Tail
	lru.Tail.Front = lru.Head
	lru.Cap = capacity
	lru.Count = 0
	lru.Map = map[int]*Node{}
	return lru
}

func (this *LRUCache) Get(key int) int {
	debug("[Get k: %v] map: %v\n", key, this.Map)
	if v, ok := this.Map[key]; !ok {
		return -1
	} else {
		this.moveToHead(v)
		debug("\t")
		printList(this.Head)
		return v.Val
	}
}

func (this *LRUCache) moveToHead(cur *Node) {
	this.remove(cur)
	this.pushToHead(cur)
}

func (this *LRUCache) pushToHead(cur *Node) {
	var temp = this.Head.Back
	this.Head.Back = cur
	cur.Front = this.Head
	cur.Back = temp
	temp.Front = cur
	this.Count++
}

func (this *LRUCache) remove(cur *Node) {
	f, b := cur.Front, cur.Back
	f.Back = b
	b.Front = f
	this.Count--
}

func (this *LRUCache) removeLRU() {
	if this.Cap <= 0 || this.Count <= 0 {
		return
	}
	key := this.Tail.Front.Key
	delete(this.Map, key)

	temp := this.Tail.Front.Front
	temp.Back = this.Tail
	this.Tail.Front = temp
	this.Count--
}

func (this *LRUCache) Put(key int, value int) {
	if this.Cap <= 0 {
		return
	}

	if v, ok := this.Map[key]; !ok {
		// add node to list
		n := Node{Val: value, Key: key}
		this.pushToHead(&n)
		this.Map[key] = &n
		// if overflow, remove
		if this.Count > this.Cap {
			this.removeLRU()
		}
	} else {
		v.Val = value
		this.moveToHead(v)
	}
	debug("[Put k: %v ,v: %v] ", key, value)
	printList(this.Head)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Get(1)
	obj.Put(3, 3)
	obj.Put(2, 1)
	obj.Get(3)
	obj.Get(4)
}
