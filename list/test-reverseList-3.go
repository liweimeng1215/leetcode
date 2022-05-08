// LC206. 反转链表 双指针解法
package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var d = false

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre, cur *ListNode
	cur = head
	for head != nil {
		debug("pre: [%v], cur: [%v]\n", pre, cur)
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
		head = temp
	}
	return pre
}
