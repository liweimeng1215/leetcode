// LC206. 反转链表 递归解法
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
	if head == nil || head.Next == nil {
		return head
	}
	ret := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}
