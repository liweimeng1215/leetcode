// 剑指 Offer 09. 用两个栈实现队列
// 用两个栈实现队列结构
package main

import "container/list"

type Stack struct {
	root *list.List
}

func (s Stack) IsEmpty() bool {
	return s.root.Front() == nil
}

func (s Stack) Push(v int) {
	s.root.PushFront(v)
}

func (s Stack) Pop() int {
	if s.IsEmpty() {
		panic("stack is empty!")
	}
	ans := s.root.Front().Value.(int)
	s.root.Remove(s.root.Front())
	return ans
}

func (s Stack) Peek() int {
	if s.IsEmpty() {
		panic("stack is empty!")
	}
	return s.root.Front().Value.(int)
}

type CQueue struct {
	stack1 Stack
	stack2 Stack
}

func Constructor() CQueue {
	return CQueue{stack1: Stack{root: &list.List{}}, stack2: Stack{root: &list.List{}}}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack1.IsEmpty() && this.stack2.IsEmpty() {
		return -1
	}

	if this.stack2.IsEmpty() {
		for !this.stack1.IsEmpty() {
			v := this.stack1.Pop()
			this.stack2.Push(v)
		}
		return this.stack2.Pop()
	} else {
		return this.stack2.Pop()
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
