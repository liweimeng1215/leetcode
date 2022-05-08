// LC102. 二叉树的层序遍历
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var l list.List
	l.PushBack(root)
	var ans = [][]int{}
	for l.Front() != nil {
		len := l.Len()
		temp := []int{}
		for i := 0; i < len; i++ {
			v := l.Front().Value.(*TreeNode)
			l.Remove(l.Front())
			temp = append(temp, v.Val)
			if v.Left != nil {
				l.PushBack(v.Left)
			}
			if v.Right != nil {
				l.PushBack(v.Right)
			}
		}
		ans = append(ans, temp)
	}
	return ans
}
