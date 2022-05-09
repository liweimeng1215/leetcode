// LC103. 二叉树的锯齿形层序遍历
// 根据当前level对链表进行不同向的访问以及节点插入
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "container/list"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var l list.List
	var level = -1
	l.PushBack(root)
	var ans = [][]int{}

	for l.Front() != nil {
		len := l.Len()
		level++
		temp := make([]int, len)
		for i := 0; i < len; i++ {
			var v *TreeNode
			if level%2 == 0 {
				v = l.Front().Value.(*TreeNode)
				l.Remove(l.Front())
				if v.Left != nil {
					l.PushBack(v.Left)
				}
				if v.Right != nil {
					l.PushBack(v.Right)
				}
			} else {
				v = l.Back().Value.(*TreeNode)
				l.Remove(l.Back())
				if v.Right != nil {
					l.PushFront(v.Right)
				}
				if v.Left != nil {
					l.PushFront(v.Left)
				}
			}
			temp[i] = v.Val
		}
		ans = append(ans, temp)
	}
	return ans
}
