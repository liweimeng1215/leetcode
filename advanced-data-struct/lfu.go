/* LC460. LFU 缓存
	LRUCache(3)
	Put(1, 1)
						| 		1		 |	 	2		 |		 3		 |
		[1]->[0x1]		(0x1)(1,1,1)

	Get(1)
						| 		1		 |	 	2		 |		 3		 |
		[1]->[0x1]							(0x1)(1,1,2)

	Put(1, 2)
						| 		1		 |	 	2		 |		 3		 |
		[1]->[0x1]											(0x1)(1,2,3)

	Put(2, 2)
						| 		1		 |	 	2		 |		 3		 |
		[1]->[0x1]											(0x1)(1,2,3)
		[2]->[0x2]			(0x2)(2,2,1)

	Put(3, 3)
						| 		1		 |	 	2		 |		 3		 |
		[1]->[0x1]											(0x1)(1,2,3)
		[2]->[0x2]			(0x3)(3,3,1)
		[3]->[0x3]			(0x2)(2,2,1)

	Put(4, 4)
						| 		1		 |	 	2		 |		 3		 |
		[1]->[0x1]											(0x1)(1,2,3)
		[3]->[0x3]			(0x4)(4,4,1)
		[4]->[0x4]			(0x3)(3,3,1)
**/

package main

import "fmt"

var d = false

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

type Node struct { // 双向链表节点
	Key   int
	Val   int
	Freq  int
	Front *Node
	Back  *Node
}

func (n Node) String() string {
	return fmt.Sprintf("{k: %v, v: %v, f: %v}", n.Key, n.Val, n.Freq)
}

func (n *Node) UpdateVal(val int) {
	n.Val = val
	n.Freq++
}

type BinaryList struct { // 双向链表
	Head  *Node
	Tail  *Node
	Count int
}

func NewBinaryList() *BinaryList {
	head, tail := &Node{}, &Node{}
	head.Back = tail
	tail.Front = head
	return &BinaryList{Head: head, Tail: tail}
}

func (bl BinaryList) String() string {
	var s string
	t := bl.Head
	for t != nil {
		s += fmt.Sprintf("%v --> ", t)
		t = t.Back
	}
	s += fmt.Sprintf("nil count: %v", bl.Count)
	return s
}

func (bl BinaryList) IsEmpty() bool {
	return bl.Count == 0
}

func (bl *BinaryList) Remove(cur *Node) {
	if bl.Count <= 0 {
		panic("no nodes to remove")
	}
	f := cur.Front
	b := cur.Back
	f.Back = b
	b.Front = f

	bl.Count--
}

func (bl *BinaryList) PushToHead(cur *Node) {
	temp := bl.Head.Back
	bl.Head.Back = cur
	cur.Front = bl.Head

	cur.Back = temp
	temp.Front = cur
	bl.Count++
}

type LFUCache struct {
	KeyToNode   map[int]*Node
	FreqToNodes map[int]*BinaryList

	Count    int
	Capacity int
	MinFreq  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{Capacity: capacity, KeyToNode: map[int]*Node{}, FreqToNodes: map[int]*BinaryList{}}
}

func (this *LFUCache) removeNode(node *Node) {
	delete(this.KeyToNode, node.Key)
	freq := node.Freq
	nodes := this.FreqToNodes[freq]
	nodes.Remove(node)
	if nodes.IsEmpty() {
		debug("[debug] delete freq: %v %v\n", freq, this.FreqToNodes[freq])
		delete(this.FreqToNodes, freq)
	}
	this.Count--
}

func (this *LFUCache) addNode(node *Node) {
	this.KeyToNode[node.Key] = node
	if nodes, ok := this.FreqToNodes[node.Freq]; ok {
		nodes.PushToHead(node)
	} else {
		nodes = NewBinaryList()
		nodes.PushToHead(node)
		this.FreqToNodes[node.Freq] = nodes
	}
	this.Count++
}

func (this *LFUCache) Get(key int) (ret int) {
	debug("[Get %v]\n", key)
	defer func() {
		debug("\tKeyToNode: %v\n", this.KeyToNode)
		debug("\tFreqToNodes: %v\n", this.FreqToNodes)
		debug("\tMinFreq: %v\n", this.MinFreq)
		debug("[Get ans: %v]\n", ret)
	}()
	if node, ok := this.KeyToNode[key]; !ok {
		return -1
	} else { // exist key
		this.removeNode(node)
		node.Freq++
		this.addNode(node)
		return node.Val
	}
}

func (this *LFUCache) Put(key int, value int) {
	debug("[Put k: %v, v: %v]\n", key, value)
	defer func() {
		debug("\tKeyToNode: %v\n", this.KeyToNode)
		debug("\tFreqToNodes: %v\n", this.FreqToNodes)
		debug("\tMinFreq: %v\n", this.MinFreq)
		debug("[Put ans: null]\n")
	}()
	if this.Capacity <= 0 {
		return
	}
	if node, ok := this.KeyToNode[key]; ok {
		this.removeNode(node)
		node.UpdateVal(value)
		this.addNode(node)
	} else {
		if this.Count >= this.Capacity {
			// remove LFU node
			for {
				if _, ok := this.FreqToNodes[this.MinFreq]; !ok {
					this.MinFreq++
				} else {
					break
				}
			}
			nodes := this.FreqToNodes[this.MinFreq]
			node = nodes.Tail.Front
			this.removeNode(node)
		}
		node = &Node{Key: key, Val: value, Freq: 1}
		this.addNode(node)
		this.MinFreq = 1
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(1)
	obj.Put(2, 1)
	obj.Get(2)
	obj.Put(3, 2)
	obj.Get(2)
	obj.Get(3)
}
