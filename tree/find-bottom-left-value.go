// LC513. 找树左下角的值
// 层序遍历，记录当前层第一个值
package main

import (
	"container/list"
	"errors"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
	ErrEmptyQueue = errors.New("queue is emtpy")
)

type queue struct {
	l      list.List
	length int
}

func (q *queue) Push(v *TreeNode) {
	q.l.PushBack(v)
	q.length++
}

func (q *queue) Front() (frontVal *TreeNode, err error) {
	frontVal, ok := q.l.Front().Value.(*TreeNode)
	if !ok {
		return nil, ErrEmptyQueue
	}
	return frontVal, nil
}

func (q *queue) Pop() (err error) {
	_, err = q.Front()
	if err != nil {
		return
	}
	q.l.Remove(q.l.Front())
	q.length--
	return nil
}

func (q *queue) IsEmpty() bool {
	return q.length == 0
}

func (q *queue) Len() int {
	return q.length
}

func findBottomLeftValue(root *TreeNode) int {
	var que queue
	que.Push(root)
	var ans = root.Val
	for !que.IsEmpty() {
		length := que.Len()
		frontVal, _ := que.Front()
		ans = frontVal.Val
		for i := 0; i < length; i++ {
			tempVal, _ := que.Front()
			if tempVal.Left != nil {
				que.Push(tempVal.Left)
			}
			if tempVal.Right != nil {
				que.Push(tempVal.Right)
			}
			que.Pop()
		}
	}
	return ans
}
