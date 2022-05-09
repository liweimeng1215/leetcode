// LC236. 二叉树的最近公共祖先
// 深度优先遍历查找到对应的节点，回溯处理结果
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p {
		return p
	}

	if root == q {
		return q
	}

	var leftAns, rightAns *TreeNode
	if root.Left != nil {
		leftAns = lowestCommonAncestor(root.Left, p, q)
	}
	if root.Right != nil {
		rightAns = lowestCommonAncestor(root.Right, p, q)
	}

	if leftAns != nil && rightAns != nil {
		return root
	}
	if leftAns != nil {
		return leftAns
	}
	if rightAns != nil {
		return rightAns
	}
	return nil
}
