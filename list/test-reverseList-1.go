// LC206. 反转链表 栈解法
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import (
	"container/list"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

var d = false

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var l list.List
	var t = head
	for t != nil {
		l.PushFront(t)
		debug("push [%v] to stack\n", t.Val)
		t = t.Next
	}
	var ans = l.Front().Value.(*ListNode)
	t = ans
	l.Remove(l.Front())
	for l.Front() != nil {
		v := l.Front().Value.(*ListNode)
		t.Next = v
		debug("t: [%v], v: [%v]\n", t.Val, v.Val)
		t = t.Next
		l.Remove(l.Front())
	}
	t.Next = nil
	return ans
}
