//LC2. 两数相加
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	addOne := 0
	ans := ListNode{}
	p := &ans
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		curBit := (v1 + v2 + addOne) % 10
		addOne = (v1 + v2 + addOne) / 10
		p.Next = &ListNode{Val: curBit}
		p = p.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if addOne != 0 {
		p.Next = &ListNode{Val: 1}
	}
	return ans.Next
}
