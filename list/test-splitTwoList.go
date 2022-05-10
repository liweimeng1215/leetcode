// 将链表分为奇偶链表
// [1] -> [2] -> [3] -> [4]
// [1] -> [3]
// [2] -> [4]
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Ans struct {
	First  *ListNode
	Second *ListNode
}

func Print(head *ListNode) {
	for head != nil {
		fmt.Printf("[%v] -> ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

func solution(head *ListNode) {
	ans := Ans{
		First:  nil,
		Second: nil,
	}

	var head1, head2 = ListNode{}, ListNode{}
	var t1, t2 = &head1, &head2
	var count = 0
	for head != nil {
		var temp = ListNode{}
		temp.Val = head.Val
		if count%2 == 0 {
			t1.Next = &temp
			t1 = t1.Next
		} else {
			t2.Next = &temp
			t2 = t2.Next
		}
		count++
		head = head.Next
	}
	ans.First = head1.Next
	ans.Second = head2.Next
	Print(ans.First)
	Print(ans.Second)
}

func main() {
	var n1, n2, n3, n4 = ListNode{1, nil}, ListNode{2, nil}, ListNode{3, nil}, ListNode{4, nil}
	n1.Next = &n2
	n2.Next = &n3
	n3.Next = &n4
	solution(&n1)
}
