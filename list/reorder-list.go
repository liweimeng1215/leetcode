// LC143. 重排链表
// [1]->[2]->[3]->[4]->[5]
// [1]->[2]->[3]
// [5]->[4]
// [1]->[5]->[2]->[4]->[3]
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

func printList(head *ListNode) {
	if d {
		for head != nil {
			fmt.Printf("[%v]-->", head.Val)
			head = head.Next
		}
	}
	fmt.Printf("nil")
	fmt.Printf("\n")
}

func reverse(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	ret := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil // [warn]!!
	return ret
}

func insert(node *ListNode, p *ListNode) (newP *ListNode) {
	temp := p.Next
	p.Next = node
	node.Next = temp
	return temp
}

func dequeue(head *ListNode) (node *ListNode, newHead *ListNode) {
	node = head
	newHead = head.Next
	node.Next = nil
	return node, newHead
}

func reorderList(head *ListNode) {
	i, j := head, head
	for i != nil && j != nil && j.Next != nil {
		i = i.Next
		j = j.Next.Next
	}
	if i == nil || i.Next == nil {
		return
	}

	j = i.Next
	i.Next = nil
	head2 := reverse(j)
	debug("-------split--------\n")
	printList(head)
	printList(head2)
	debug("-------split--------\n")

	i, j = head, head2
	for j != nil {
		node, newj := dequeue(j)
		i = insert(node, i)
		j = newj
		printList(i)
		printList(j)
	}
}
