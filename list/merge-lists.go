// LC23. 合并K个升序链表
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func merge(l1, l2 *ListNode) *ListNode {
	i, j := l1, l2
	ans := ListNode{}
	t := &ans
	for j != nil && i != nil {
		if j.Val < i.Val {
			t.Next = j
			j = j.Next
			t = t.Next
		} else {
			t.Next = i
			i = i.Next
			t = t.Next
		}
	}
	if j == nil {
		t.Next = i
		return ans.Next
	}
	t.Next = j
	return ans.Next
}

func mergeRange(list []*ListNode, lo, hi int) *ListNode {
	if lo == hi {
		return list[lo]
	}
	mi := lo + (hi-lo)/2
	l1 := mergeRange(list, lo, mi)
	l2 := mergeRange(list, mi+1, hi)
	return merge(l1, l2)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	ans := mergeRange(lists, 0, len(lists)-1)
	return ans
}
